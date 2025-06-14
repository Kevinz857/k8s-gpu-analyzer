package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Kevinz857/k8s-gpu-analyzer/internal/k8s"
	"github.com/Kevinz857/k8s-gpu-analyzer/internal/monitor"
	"github.com/spf13/cobra"
)

var (
	nodeLabels []string
	namespaces []string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "k8s-gpu-analyzer",
		Short: "Kubernetes GPU resource analyzer",
		Long:  `A tool to analyze GPU resource usage across Kubernetes nodes and pods`,
		Run:   runAnalyzer,
	}

	// Add flags
	rootCmd.Flags().StringSliceVarP(&nodeLabels, "node-labels", "l", []string{"gpu=true"},
		"Node labels to filter GPU nodes (format: key=value,key2=value2)")
	rootCmd.Flags().StringSliceVarP(&namespaces, "namespaces", "n", []string{"default"},
		"Namespaces to search for pods (comma-separated)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runAnalyzer(cmd *cobra.Command, args []string) {
	// Create Kubernetes client
	clientset, err := k8s.CreateClient()
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	// Parse node labels
	nodeLabelMap := make(map[string]string)
	for _, label := range nodeLabels {
		if strings.Contains(label, "=") {
			parts := strings.SplitN(label, "=", 2)
			nodeLabelMap[parts[0]] = parts[1]
		} else {
			log.Fatalf("Invalid node label format: %s (expected key=value)", label)
		}
	}

	// Create GPU analyzer
	gpuAnalyzer := monitor.NewGPUAnalyzer(clientset)

	// Get GPU node information
	gpuNodes, err := gpuAnalyzer.GetGPUNodeInfo(nodeLabelMap, namespaces)
	if err != nil {
		log.Fatalf("Failed to get GPU node information: %v", err)
	}

	// Print results
	monitor.PrintGPUNodeInfo(gpuNodes)
}
