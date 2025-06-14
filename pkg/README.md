# Public Packages

本目录包含可以被外部应用程序使用的库代码。这些包是项目的公共 API，其他项目可以安全地导入和使用。

## 包结构

### types/
共享数据类型定义。

- `types.go`: 定义了项目中使用的核心数据结构，如 `GPUNodeInfo`

## 使用示例

```go
import "github.com/Kevin857/k8s-gpu-analyzer/pkg/types"

// 使用公共类型
var nodeInfo types.GPUNodeInfo
nodeInfo.NodeName = "gpu-node-001"
nodeInfo.GPUPodCount = 5
```

## 设计原则

- **稳定的 API**: 公共包的接口应该保持向后兼容
- **最小化依赖**: pkg 包应该尽量减少外部依赖
- **清晰的文档**: 所有公共接口都应该有清晰的文档说明
- **类型安全**: 使用强类型定义，避免使用 interface{} 或 any
