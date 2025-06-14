package monitor

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/Kevin857/k8s-gpu-analyzer/pkg/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GPUAnalyzer represents a GPU analysis service
type GPUAnalyzer struct {
	clientset *kubernetes.Clientset
}

// NewGPUAnalyzer creates a new GPU analyzer instance
func NewGPUAnalyzer(clientset *kubernetes.Clientset) *GPUAnalyzer {
	return &GPUAnalyzer{
		clientset: clientset,
	}
}

// GetGPUNodeInfo retrieves information and statistics for all GPU nodes
func (a *GPUAnalyzer) GetGPUNodeInfo(nodeLabels map[string]string, namespaces []string) ([]types.GPUNodeInfo, error) {
	ctx := context.Background()

	// Build node selector for filtering GPU nodes
	nodeSelector := ""
	if len(nodeLabels) > 0 {
		var selectors []string
		for key, value := range nodeLabels {
			selectors = append(selectors, fmt.Sprintf("%s=%s", key, value))
		}
		nodeSelector = strings.Join(selectors, ",")
	}

	// Get all nodes with optional label selector
	nodes, err := a.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{
		LabelSelector: nodeSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get node list: %v", err)
	}

	// Determine which namespaces to query
	if len(namespaces) == 0 {
		namespaces = []string{"default"}
	}

	// Filter GPU nodes first
	var gpuNodes []corev1.Node
	var gpuNodeNames []string
	for _, node := range nodes.Items {
		if a.isGPUNode(&node, nodeLabels) {
			gpuNodes = append(gpuNodes, node)
			gpuNodeNames = append(gpuNodeNames, node.Name)
		}
	}

	if len(gpuNodes) == 0 {
		return []types.GPUNodeInfo{}, nil
	}

	// Build field selector for all GPU nodes to minimize API calls
	// This is much more efficient than querying each node separately
	nodeFieldSelector := ""
	if len(gpuNodeNames) == 1 {
		nodeFieldSelector = fmt.Sprintf("spec.nodeName=%s", gpuNodeNames[0])
	} else {
		// For multiple nodes, we'll use a more efficient approach
		// Unfortunately, Kubernetes doesn't support "in" operator for field selectors
		// So we'll query each namespace once for all nodes, then filter client-side
		nodeFieldSelector = "" // Query all pods in namespace, filter client-side
	}

	// Create a map to store pods for each node
	nodePods := make(map[string][]corev1.Pod)
	for _, nodeName := range gpuNodeNames {
		nodePods[nodeName] = []corev1.Pod{}
	}

	// Query pods more efficiently
	for _, namespace := range namespaces {
		var pods *corev1.PodList

		if len(gpuNodeNames) == 1 {
			// For single node, use field selector for maximum efficiency
			pods, err = a.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
				FieldSelector: nodeFieldSelector,
			})
		} else {
			// For multiple nodes, query all pods in namespace once
			// This reduces API calls from N*M to M (where N=nodes, M=namespaces)
			pods, err = a.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
		}

		if err != nil {
			return nil, fmt.Errorf("failed to get pod list in namespace %s: %v", namespace, err)
		}

		// Filter pods for GPU nodes
		for _, pod := range pods.Items {
			// Only include pods on GPU nodes
			if contains(gpuNodeNames, pod.Spec.NodeName) {
				nodePods[pod.Spec.NodeName] = append(nodePods[pod.Spec.NodeName], pod)
			}
		}
	}

	var gpuNodeInfos []types.GPUNodeInfo

	// Process each GPU node
	for _, node := range gpuNodes {
		nodeInfo := types.GPUNodeInfo{
			NodeName: node.Name,
		}

		// Get total GPU capacity of the node
		if gpuQuantity, exists := node.Status.Capacity["nvidia.com/gpu"]; exists {
			nodeInfo.NodeGPUTotal = gpuQuantity.Value()
		}

		// Count GPU pods on this node
		gpuPodCount := 0
		var totalGPURequest int64 = 0

		for _, pod := range nodePods[node.Name] {
			// Skip completed or failed pods
			if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
				continue
			}

			// Check if pod uses GPU
			podGPURequest := int64(0)
			for _, container := range pod.Spec.Containers {
				if gpuQuantity, exists := container.Resources.Requests["nvidia.com/gpu"]; exists {
					podGPURequest += gpuQuantity.Value()
				}
				if gpuQuantity, exists := container.Resources.Limits["nvidia.com/gpu"]; exists {
					if requestQuantity, requestExists := container.Resources.Requests["nvidia.com/gpu"]; !requestExists || requestQuantity.IsZero() {
						podGPURequest += gpuQuantity.Value()
					}
				}
			}

			if podGPURequest > 0 {
				gpuPodCount++
				totalGPURequest += podGPURequest
			}
		}

		nodeInfo.GPUPodCount = gpuPodCount
		nodeInfo.NodeGPURequest = totalGPURequest

		// Calculate GPU usage percentage
		if nodeInfo.NodeGPUTotal > 0 {
			nodeInfo.NodeGPURequestPercent = float64(nodeInfo.NodeGPURequest) / float64(nodeInfo.NodeGPUTotal) * 100
		}

		gpuNodeInfos = append(gpuNodeInfos, nodeInfo)
	}

	// Sort by node name
	sort.Slice(gpuNodeInfos, func(i, j int) bool {
		return gpuNodeInfos[i].NodeName < gpuNodeInfos[j].NodeName
	})

	return gpuNodeInfos, nil
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// isGPUNode checks whether a node is a GPU node
func (a *GPUAnalyzer) isGPUNode(node *corev1.Node, nodeLabels map[string]string) bool {
	// If specific node labels are provided, check if the node matches
	if len(nodeLabels) > 0 {
		for key, value := range nodeLabels {
			if nodeValue, exists := node.Labels[key]; !exists || nodeValue != value {
				return false
			}
		}
		return true
	}

	// Check if node has GPU resources
	if _, exists := node.Status.Capacity["nvidia.com/gpu"]; exists {
		return true
	}

	// Check node labels for GPU-related keywords
	for key := range node.Labels {
		if strings.Contains(strings.ToLower(key), "gpu") ||
			strings.Contains(strings.ToLower(key), "nvidia") ||
			strings.Contains(strings.ToLower(key), "accelerator") {
			return true
		}
	}

	// Check node name for GPU-related keywords
	nodeName := strings.ToLower(node.Name)
	if strings.Contains(nodeName, "gpu") ||
		strings.Contains(nodeName, "nvidia") {
		return true
	}

	return false
}
