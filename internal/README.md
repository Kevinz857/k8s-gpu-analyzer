# Internal Packages

本目录包含项目的私有应用程序和库代码。根据 Go 的 `internal` 包规范，其他项目无法导入这些包。

## 包结构

### k8s/
Kubernetes 客户端相关功能。

- `client.go`: 提供创建 Kubernetes clientset 的功能，支持 out-of-cluster 访问方式

### monitor/
GPU 监控核心逻辑。

- `gpu_monitor.go`: GPU 监控器的主要实现，包含节点发现、Pod 统计等功能
- `printer.go`: 输出格式化功能，负责将监控数据以表格形式输出

## 设计原则

- **单一职责**: 每个包只负责一个特定的功能领域
- **依赖注入**: 通过构造器注入依赖，便于测试和扩展
- **错误处理**: 所有公共函数都有适当的错误处理和返回
- **可测试性**: 代码结构便于单元测试
