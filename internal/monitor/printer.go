package monitor

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Kevin857/k8s-gpu-analyzer/pkg/types"
)

// PrintGPUNodeInfo prints GPU node information in table format
func PrintGPUNodeInfo(gpuNodes []types.GPUNodeInfo) {
	if len(gpuNodes) == 0 {
		fmt.Println("No GPU nodes found")
		return
	}

	// Create tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	// Print table header
	fmt.Fprintln(w, "GPU-Node-Name\tGPUPodCount\tNodeGPURequest\tNodeGPURequestPercent\tNodeGPUTotal")
	fmt.Fprintln(w, "-------------\t-----------\t--------------\t--------------------\t------------")

	// Print information for each GPU node
	for _, nodeInfo := range gpuNodes {
		fmt.Fprintf(w, "%s\t%d\t%d\t%.2f%%\t%d\n",
			nodeInfo.NodeName,
			nodeInfo.GPUPodCount,
			nodeInfo.NodeGPURequest,
			nodeInfo.NodeGPURequestPercent,
			nodeInfo.NodeGPUTotal,
		)
	}

	// Calculate totals
	totalPods := 0
	totalRequest := int64(0)
	totalCapacity := int64(0)

	for _, nodeInfo := range gpuNodes {
		totalPods += nodeInfo.GPUPodCount
		totalRequest += nodeInfo.NodeGPURequest
		totalCapacity += nodeInfo.NodeGPUTotal
	}

	// Print separator line
	fmt.Fprintln(w, "-------------\t-----------\t--------------\t--------------------\t------------")

	// Print totals
	totalPercent := float64(0)
	if totalCapacity > 0 {
		totalPercent = float64(totalRequest) / float64(totalCapacity) * 100
	}

	fmt.Fprintf(w, "TOTAL\t%d\t%d\t%.2f%%\t%d\n",
		totalPods,
		totalRequest,
		totalPercent,
		totalCapacity,
	)

	// Flush output
	w.Flush()

	// Print summary statistics
	fmt.Printf("\nSummary:\n")
	fmt.Printf("Total GPU nodes: %d\n", len(gpuNodes))
	fmt.Printf("Total GPU pods: %d\n", totalPods)
	fmt.Printf("Total GPU requests: %d\n", totalRequest)
	fmt.Printf("Total GPU capacity: %d\n", totalCapacity)
	fmt.Printf("Overall GPU utilization: %.2f%%\n", totalPercent)
}
