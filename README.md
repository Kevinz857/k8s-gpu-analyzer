# K8s GPU Analyzer

[![Go Version](https://img.shields.io/github/go-mod/go-version/Kevinz857/k8s-gpu-analyzer)](https://github.com/Kevinz857/k8s-gpu-analyzer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CI](https://github.com/Kevinz857/k8s-gpu-analyzer/workflows/CI/badge.svg)](https://github.com/Kevinz857/k8s-gpu-analyzer/actions)
[![Release](https://img.shields.io/github/v/release/Kevinz857/k8s-gpu-analyzer)](https://github.com/Kevinz857/k8s-gpu-analyzer/releases)

🚀 **A high-performance Golang tool for analyzing GPU resource allocation and utilization in Kubernetes clusters**

*Reduce API server load by up to 75% while getting comprehensive GPU usage insights*

**⭐ If this project helps you, please give it a star! ⭐**

## Features

- Automatically discover GPU nodes in the cluster
- Count GPU pods on each node
- Calculate GPU resource requests and utilization rates
- Display statistics in a clear table format
- Support out-of-cluster access to Kubernetes
- Filter GPU nodes by custom labels
- Query pods from specific namespaces to reduce apiserver load

## Prerequisites

- Go 1.21 or higher
- Valid kubeconfig file (usually located at `~/.kube/config`)
- Access permissions to the Kubernetes cluster

## Installation and Build

### Option 1: Download Pre-built Binaries (Recommended)

Download the latest release from [GitHub Releases](https://github.com/Kevinz857/k8s-gpu-analyzer/releases):

**Linux:**
```bash
# AMD64
wget https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-linux-amd64
chmod +x k8s-gpu-analyzer-linux-amd64
./k8s-gpu-analyzer-linux-amd64

# ARM64
wget https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-linux-arm64
chmod +x k8s-gpu-analyzer-linux-arm64
./k8s-gpu-analyzer-linux-arm64
```

**macOS:**
```bash
# Intel
wget https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-darwin-amd64
chmod +x k8s-gpu-analyzer-darwin-amd64
./k8s-gpu-analyzer-darwin-amd64

# Apple Silicon
wget https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-darwin-arm64
chmod +x k8s-gpu-analyzer-darwin-arm64
./k8s-gpu-analyzer-darwin-arm64
```

**Windows:**
```powershell
# AMD64
Invoke-WebRequest -Uri https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-windows-amd64.exe -OutFile k8s-gpu-analyzer.exe
.\k8s-gpu-analyzer.exe

# ARM64
Invoke-WebRequest -Uri https://github.com/Kevinz857/k8s-gpu-analyzer/releases/latest/download/k8s-gpu-analyzer-windows-arm64.exe -OutFile k8s-gpu-analyzer.exe
.\k8s-gpu-analyzer.exe
```

### Option 2: Build from Source

```bash
# Clone the project
git clone https://github.com/Kevinz857/k8s-gpu-analyzer.git
cd k8s-gpu-analyzer

# Download dependencies
go mod tidy

# Build the project
make build

# Run the program
./bin/k8s-gpu-analyzer
```

## Usage

### Basic Usage

```bash
# Use default settings (gpu=true label, default namespace)
./bin/k8s-gpu-analyzer

# Specify custom node labels
./bin/k8s-gpu-analyzer --node-labels "gpu=true,instance-type=gpu"

# Specify custom namespaces
./bin/k8s-gpu-analyzer --namespaces "default,kube-system,gpu-namespace"

# Combine both options
./bin/k8s-gpu-analyzer --node-labels "gpu=true" --namespaces "default,production"
```

### Command Line Options

```bash
./bin/k8s-gpu-analyzer --help
```

Available flags:
- `-l, --node-labels`: Node labels to filter GPU nodes (format: key=value,key2=value2) (default: gpu=true)
- `-n, --namespaces`: Namespaces to search for pods (comma-separated) (default: default)
- `-h, --help`: Show help information

## Configuration

The program looks for kubeconfig files in the following order:

1. Path specified by the `KUBECONFIG` environment variable
2. Default path `~/.kube/config`

To specify a custom kubeconfig file:

```bash
export KUBECONFIG=/path/to/your/kubeconfig
./bin/k8s-gpu-analyzer
```

## Output Format

The program outputs a table in the following format:

```
GPU-Node-Name          GPUPodCount    NodeGPURequest    NodeGPURequestPercent    NodeGPUTotal
-------------          -----------    --------------    --------------------     ------------
gpu-node-001                     3                 6                75.00%                8
gpu-node-002                     2                 4                50.00%                8
gpu-node-003                     1                 2                25.00%                8
-------------          -----------    --------------    --------------------     ------------
TOTAL                            6                12                50.00%               24

Summary:
Total GPU nodes: 3
Total GPU pods: 6
Total GPU requests: 12
Total GPU capacity: 24
Overall GPU utilization: 50.00%
```

## Field Descriptions

- **GPU-Node-Name**: GPU node name
- **GPUPodCount**: Number of GPU-using pods on the node
- **NodeGPURequest**: Total GPU resource requests on the node
- **NodeGPURequestPercent**: GPU resource utilization percentage
- **NodeGPUTotal**: Total GPU capacity of the node

## GPU Node Identification Rules

The program identifies GPU nodes using the following rules:

1. **Custom Labels**: If `--node-labels` is specified, nodes must match all specified labels
2. **GPU Resources**: Node resources contain `nvidia.com/gpu`
3. **Label Keywords**: Node labels contain `gpu`, `nvidia`, or `accelerator` keywords
4. **Name Keywords**: Node names contain `gpu` or `nvidia` keywords

## Performance Optimizations

To reduce load on the apiserver, especially in large clusters, the tool uses several optimization strategies:

### API Call Optimization Strategy

The tool intelligently minimizes API calls based on the number of GPU nodes:

1. **Single GPU Node**: Uses field selectors for maximum efficiency
   - API calls: `1 nodes.list() + M pods.list()` (where M = number of namespaces)
   - Example: 1 node + 3 namespaces = 4 API calls total

2. **Multiple GPU Nodes**: Batch queries to reduce API calls
   - API calls: `1 nodes.list() + M pods.list()` (same as single node!)
   - Filters pods client-side to avoid N×M API calls
   - Example: 5 nodes + 3 namespaces = 4 API calls (not 16!)

### Optimization Techniques

1. **Node Label Filtering**: Use `--node-labels` to filter nodes at the Kubernetes API level
2. **Namespace Filtering**: By default, only queries the `default` namespace
3. **Smart Pod Querying**: 
   - Single node: Uses `spec.nodeName=<node>` field selector
   - Multiple nodes: Queries all pods per namespace once, filters client-side
4. **Early GPU Node Detection**: Filters GPU nodes before any pod queries

### Alternative Approaches Considered

We evaluated several methods for gathering GPU usage information:

| Method | API Calls | Pros | Cons |
|--------|-----------|------|------|
| **Current Optimized** | `1 + M` | Minimal API calls, accurate | Some client-side filtering |
| Original Per-Node | `1 + (N × M)` | Simple logic | High API server load |
| Metrics API | `2-3` | Very low API calls | Requires metrics-server, no pod count |
| Node Status Only | `1` | Minimal load | No pod-level details, less accurate |
| Event-based | `2-3` | Low API calls | Unreliable, events expire |

**Legend**: N = GPU nodes, M = namespaces

### Performance Comparison

For a cluster with 5 GPU nodes and 3 namespaces:

- **Naive approach**: 1 + (5 × 3) = **16 API calls**
- **Our optimized approach**: 1 + 3 = **4 API calls** (75% reduction)

The optimization becomes more significant as the number of GPU nodes increases.

## Notes

- Only counts pods in Running and Pending states
- Supports checking both requests and limits for GPU resources in pods
- If a pod only sets limits without requests, uses the limits value
- All GPU-related resources are based on the `nvidia.com/gpu` resource type

## Troubleshooting

### Common Errors

1. **Failed to connect to Kubernetes cluster**
   - Check if kubeconfig file exists and is valid
   - Verify network connectivity and cluster access permissions

2. **No GPU nodes found**
   - Confirm that GPU nodes actually exist in the cluster
   - Check if nodes are properly configured with GPU resources
   - Verify that node labels match the specified `--node-labels`

3. **Permission errors**
   - Ensure the kubeconfig user has sufficient permissions to access node and pod information

## Development

The project follows the [Go Standard Project Layout](https://github.com/golang-standards/project-layout):

```
.
├── cmd/
│   └── k8s-gpu-analyzer/   # Main application
│       └── main.go
├── internal/              # Private application and library code
│   ├── k8s/              # Kubernetes client
│   │   └── client.go
│   └── monitor/          # GPU analysis core logic
│       ├── gpu_analyzer.go
│       └── printer.go
├── pkg/                  # Library code that can be used by external applications
│   └── types/           # Public type definitions
│       └── types.go
├── go.mod               # Go module definition
├── Makefile            # Build scripts
├── README.md           # Documentation
└── .gitignore          # Git ignore file
```

Main modules:

- `cmd/k8s-gpu-analyzer`: Main application entry point
- `internal/k8s`: Kubernetes client creation and configuration
- `internal/monitor`: GPU monitoring core logic and output formatting
- `pkg/types`: Shared data type definitions

## License

MIT License
