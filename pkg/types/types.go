package types

// GPUNodeInfo stores statistical information for GPU nodes
type GPUNodeInfo struct {
	NodeName              string  // Node name
	GPUPodCount           int     // Number of GPU pods on the node
	NodeGPURequest        int64   // Total GPU requests on the node
	NodeGPUTotal          int64   // Total GPU capacity of the node
	NodeGPURequestPercent float64 // GPU request percentage
}
