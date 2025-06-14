---
name: Bug report
about: Create a report to help us improve
title: '[BUG] '
labels: bug
assignees: ''

---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Run command '...'
2. With cluster configuration '...'
3. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Environment:**
- Kubernetes version: [e.g. v1.28.0]
- GPU node configuration: [e.g. NVIDIA Tesla V100, 8 GPUs per node]
- Number of GPU nodes: [e.g. 5]
- Operating System: [e.g. Ubuntu 20.04]
- k8s-gpu-analyzer version: [e.g. v1.0.0]

**Command used:**
```bash
./k8s-gpu-analyzer --node-labels "gpu=true" --namespaces "default,kube-system"
```

**Error output:**
```
Paste the error output here
```

**Additional context**
Add any other context about the problem here.
