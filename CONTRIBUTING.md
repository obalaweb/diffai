# Contributing to DiffAI

Thank you for your interest in contributing to DiffAI! This document provides guidelines and information for contributors.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- Python 3.11 or later
- Git
- Docker (optional, for containerized development)

### Development Setup

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/your-username/diffai.git
   cd diffai
   ```

2. **Set up the development environment**
   ```bash
   make setup
   ```

3. **Start the AI service for development**
   ```bash
   make dev
   ```

4. **Build the CLI**
   ```bash
   make build
   ```

## 🏗️ Architecture Overview

DiffAI follows a modular architecture:

- **CLI Layer (Go)**: Command-line interface using Cobra
- **Git Integration (Go)**: Git operations using go-git
- **AI Service (Python)**: FastAPI microservice for AI operations
- **Configuration**: YAML-based configuration with environment variable support

## 📝 Development Workflow

### Making Changes

1. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Follow Go and Python coding standards
   - Add tests for new functionality
   - Update documentation as needed

3. **Run tests**
   ```bash
   make test
   ```

4. **Format code**
   ```bash
   make fmt
   ```

5. **Run linters**
   ```bash
   make lint
   ```

### Testing

- **Go tests**: `go test ./...`
- **Python tests**: `cd ai-service && python -m pytest`
- **Integration tests**: Use the test fixtures in `tests/fixtures/`

### Code Style

#### Go
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused

#### Python
- Follow PEP 8 style guidelines
- Use type hints where appropriate
- Add docstrings for functions and classes
- Use meaningful variable names

## 🐛 Bug Reports

When reporting bugs, please include:

1. **Environment information**
   - OS and version
   - Go version
   - Python version
   - DiffAI version

2. **Steps to reproduce**
   - Clear, numbered steps
   - Expected vs actual behavior

3. **Additional context**
   - Error messages
   - Log files
   - Configuration files (with sensitive data removed)

## ✨ Feature Requests

When requesting features, please:

1. **Check existing issues** to avoid duplicates
2. **Describe the use case** and why it would be valuable
3. **Provide examples** of how the feature would work
4. **Consider implementation complexity** and alternatives

## 📋 Pull Request Process

1. **Update documentation** for any new features
2. **Add tests** for new functionality
3. **Ensure all tests pass**
4. **Update CHANGELOG.md** with your changes
5. **Request review** from maintainers

### PR Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass locally
- [ ] New tests added for new functionality
- [ ] Manual testing completed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
```

## 🏷️ Release Process

Releases are managed by maintainers:

1. **Version bump** in `cmd/diffai/main.go`
2. **Update CHANGELOG.md**
3. **Create release tag**
4. **Build and publish** binaries
5. **Update documentation**

## 🤝 Community Guidelines

### Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow
- Follow the golden rule

### Communication

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and ideas
- **Pull Requests**: Code contributions and reviews

## 🛠️ Advanced Development

### Adding New AI Providers

1. **Create provider in `internal/ai/providers/`**
2. **Implement the provider interface**
3. **Add configuration options**
4. **Update documentation**

### Adding New Commands

1. **Create command in `internal/cli/`**
2. **Add to main command structure**
3. **Implement Git integration if needed**
4. **Add tests and documentation**

### Docker Development

```bash
# Build and run with Docker Compose
make compose-up

# Run specific service
docker-compose up ai-service

# View logs
make logs
```

## 📚 Resources

- [Go Documentation](https://golang.org/doc/)
- [Cobra CLI Framework](https://cobra.dev/)
- [FastAPI Documentation](https://fastapi.tiangolo.com/)
- [Conventional Commits](https://www.conventionalcommits.org/)

## ❓ Questions?

If you have questions about contributing:

1. Check existing [GitHub Discussions](https://github.com/diffai/diffai/discussions)
2. Open a new discussion for general questions
3. Open an issue for specific problems

Thank you for contributing to DiffAI! 🎉
