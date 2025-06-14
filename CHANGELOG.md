# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2024-12-14

### Added
- Initial release of K8s GPU Analyzer
- GPU node discovery and analysis
- Pod GPU resource usage tracking
- Command-line interface with customizable options
- Node label filtering (`--node-labels`)
- Namespace filtering (`--namespaces`)
- Optimized API queries to reduce apiserver load
- Comprehensive documentation and examples

### Features
- **Smart GPU Node Detection**: Automatically discovers GPU nodes using multiple strategies
- **Performance Optimized**: Reduces API calls by up to 75% compared to naive approaches
- **Flexible Filtering**: Support for custom node labels and namespace selection
- **Clear Output**: Table format with detailed statistics and summaries
- **Production Ready**: Designed for large-scale Kubernetes clusters

### Technical Highlights
- **API Optimization**: Intelligent query batching (1 + M API calls vs 1 + N×M)
- **Field Selectors**: Efficient pod filtering using Kubernetes field selectors
- **Client-side Filtering**: Minimizes network overhead for multi-node scenarios
- **Resource Accuracy**: Supports both requests and limits for GPU resources

[Unreleased]: https://github.com/Kevin857/k8s-gpu-analyzer/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/Kevin857/k8s-gpu-analyzer/releases/tag/v1.0.0
