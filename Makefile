.PHONY: help build run clean tidy test

# 默认目标
help:
	@echo "K8s GPU Analyzer - Kubernetes GPU 节点资源分析工具"
	@echo ""
	@echo "可用命令:"
	@echo "  build   - 编译项目"
	@echo "  run     - 运行程序"
	@echo "  clean   - 清理构建文件"
	@echo "  tidy    - 整理依赖"
	@echo "  test    - 运行测试"
	@echo "  help    - 显示帮助信息"

# 编译项目
build:
	@echo "正在编译 k8s-gpu-analyzer..."
	go build -o k8s-gpu-analyzer ./cmd/k8s-gpu-analyzer
	@echo "编译完成！"

# 运行程序
run: build
	@echo "运行 K8s GPU 分析工具..."
	./k8s-gpu-analyzer

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -f k8s-gpu-analyzer
	go clean
	@echo "清理完成！"

# 整理依赖
tidy:
	@echo "整理 Go 模块依赖..."
	go mod tidy
	@echo "依赖整理完成！"

# 运行测试（目前没有测试文件）
test:
	@echo "运行测试..."
	go test ./...
	@echo "测试完成！"

# 安装到系统路径（可选）
install: build
	@echo "安装 k8s-gpu-analyzer 到 /usr/local/bin..."
	sudo cp k8s-gpu-analyzer /usr/local/bin/
	@echo "安装完成！现在可以在任何地方运行 'k8s-gpu-analyzer' 命令"

# 检查代码格式
fmt:
	@echo "格式化代码..."
	go fmt ./...
	@echo "代码格式化完成！"

# 代码检查
vet:
	@echo "运行代码检查..."
	go vet ./...
	@echo "代码检查完成！"
