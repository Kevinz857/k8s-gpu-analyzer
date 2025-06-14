# Contributing to K8s GPU Analyzer

Thank you for your interest in contributing to K8s GPU Analyzer! This document outlines the process for contributing to this project.

## How to Contribute

### Reporting Issues
- Use GitHub Issues to report bugs or request features
- Provide detailed information including:
  - Kubernetes version
  - GPU node configuration
  - Steps to reproduce the issue
  - Expected vs actual behavior

### Submitting Changes
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Setup
```bash
# Clone your fork
git clone https://github.com/yourusername/k8s-gpu-analyzer.git
cd k8s-gpu-analyzer

# Install dependencies
go mod tidy

# Build the project
make build

# Run tests
go test ./...
```

## Code Style
- Follow Go best practices and conventions
- Use `gofmt` to format your code
- Add comments for public functions and complex logic
- Keep functions focused and testable

## Feature Requests
We welcome feature requests! Please:
- Check existing issues first
- Describe the use case and expected behavior
- Consider contributing the implementation

## Questions?
Feel free to open an issue for any questions about contributing.
