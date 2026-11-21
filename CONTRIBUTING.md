# Contributing to Chamego

Thank you for your interest in contributing to Chamego! This document outlines the guidelines and processes for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Submitting Pull Requests](#submitting-pull-requests)
- [Coding Standards](#coding-standards)
- [Issue Guidelines](#issue-guidelines)

## Code of Conduct

Please note that this project is released with a [Code of Conduct](./CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- A code editor (VS Code, GoLand, etc.)
- Git

### Setup

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/yourusername/chamego.git
   cd chamego
   ```
3. Create a branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Workflow

### Branching Strategy

- Use feature branches based on `main`
- Name branches using `<type>/<description>` format:
  - `feature/<name>` for new features
  - `bugfix/<name>` for bug fixes
  - `docs/<name>` for documentation changes
  - `refactor/<name>` for code refactoring
  - `test/<name>` for test additions

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Test changes
- `chore`: Build process or auxiliary tool changes

## Submitting Pull Requests

### Before Submitting

1. Ensure your code follows the project's coding standards
2. All tests pass:
   ```bash
   go test ./...
   ```
3. Run linters:
   ```bash
   golangci-lint run
   ```
4. Update documentation if needed
5. Add relevant issue references in your PR description

### Pull Request Template

When creating a PR, include:

- [ ] Description of changes
- [ ] Related issue number (if applicable)
- [ ] Screenshots (for UI changes)
- [ ] Testing performed
- [ ] Breaking changes (if any)

## Coding Standards

### Go Specific

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `go fmt` for code formatting
- Write meaningful comments and docstrings
- Keep functions small and focused
- Use table-driven tests for unit tests

### Code Style

```go
// Package name provides package description.
package name

import (
    // Import standard library packages in order:
    // 1. context, encoding, flag, fmt, io, log, net, os, path/filepath, reflect, sort, strconv, sync, testing
    // 2. third-party packages
    // 3. local packages
)

// FunctionName does something meaningful.
func FunctionName(arg string) (result string, err error) {
    // Implementation
}
```

## Issue Guidelines

### When Creating an Issue

1. **Use a clear and descriptive title**
2. **Provide as much context as possible**
3. **Include steps to reproduce** (for bugs)
4. **Search existing issues first** to avoid duplicates
5. **Use appropriate labels** when possible

### Issue Types

- **Bug Reports**: Include expected vs actual behavior, steps to reproduce, and environment details
- **Feature Requests**: Explain the use case and benefits
- **Questions**: Use the discussion template or ask in community channels

## Review Process

1. PRs are reviewed by maintainers
2. Reviewers may request changes or improvements
3. All feedback should be addressed before merging
4. Maintainers have final say on what gets merged

## Getting Help

- Check existing issues and documentation
- Ask questions in the project's discussion channels
- Reach out to maintainers for clarification

## License

By contributing, you agree that your contributions will be licensed under the [project's LICENSE](./LICENSE) file.

---

Thank you for contributing to Chamego! 🚀
