# DiffAI 🧠✨

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Python Version](https://img.shields.io/badge/Python-3.11+-3776AB?style=for-the-badge&logo=python)](https://python.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![GitHub Issues](https://img.shields.io/github/issues/obalaweb/diffai?style=for-the-badge)](https://github.com/obalaweb/diffai/issues)
[![GitHub Stars](https://img.shields.io/github/stars/obalaweb/diffai?style=for-the-badge)](https://github.com/obalaweb/diffai/stargazers)

**AI-powered Git assistant that understands your diffs.**

DiffAI is an open-source tool that helps developers create better commits, generate meaningful PR summaries, and automate changelogs — all powered by AI that *actually understands code changes*.

Unlike tools that only automate commit messages, DiffAI analyzes your diffs and creates intelligent summaries that make your Git workflow more productive and professional.

---

## 🎯 Why DiffAI?

### The Problem
- **Poor commit messages**: Most commits are vague ("fix bug", "update code")
- **Tedious PR descriptions**: Writing clear PR summaries is time-consuming
- **Changelog maintenance**: Keeping changelogs updated is painful but crucial
- **Limited AI tools**: Existing solutions are closed-source or lack deep code understanding

### The Solution
DiffAI fills this gap with an **open-source, AI-first, developer-centric Git companion** that:
- 🤖 **Understands your code changes** and generates meaningful commit messages
- 📑 **Analyzes PR diffs** to create comprehensive summaries
- 📝 **Groups commits intelligently** for professional changelogs
- 🔌 **Integrates seamlessly** with your existing Git workflow

---

## 🚀 Features

### ✅ **AI-Powered Commit Messages**
```bash
git add .
git diffai commit
# ✨ Generates: "feat: add user authentication with JWT tokens"
```

### ✅ **Intelligent PR Summaries**
```bash
git diffai pr 123
# 📋 Creates comprehensive PR overview with impact analysis
```

### ✅ **Automated Changelogs**
```bash
git diffai changelog
# 📝 Generates structured changelog grouped by features & fixes
```

### ✅ **Multiple AI Providers**
- **OpenAI GPT-4** - Most advanced reasoning
- **Anthropic Claude** - Excellent code understanding
- **Local Models** - Privacy-focused with Ollama

### ✅ **Developer-Friendly**
- 🔌 **Git CLI Integration** - Works as `git diffai` command
- ⚙️ **Flexible Configuration** - YAML config with environment overrides
- 🐳 **Docker Support** - Easy deployment and scaling
- 🌍 **Cross-Platform** - Linux, macOS, Windows

---

## 📦 Installation

### Option 1: Download Binary (Recommended)
```bash
# Download latest release
curl -L https://github.com/obalaweb/diffai/releases/latest/download/diffai-linux-amd64 -o diffai
chmod +x diffai
sudo mv diffai /usr/local/bin/
```

### Option 2: Build from Source
```bash
git clone https://github.com/obalaweb/diffai.git
cd diffai
make build
sudo make install
```

### Option 3: Docker
```bash
docker run -it --rm -v $(pwd):/workspace obalaweb/diffai:latest
```

---

## ⚡ Quick Start

### 1. **Initialize Configuration**
```bash
git diffai config init
```

### 2. **Set AI Provider**
```bash
# For OpenAI
export OPENAI_API_KEY="your-api-key"

# For Anthropic
export ANTHROPIC_API_KEY="your-api-key"

# For Local AI (Ollama)
export LOCAL_AI_BASE_URL="http://localhost:11434"
```

### 3. **Generate Commit Message**
```bash
git add .
git diffai commit
```

### 4. **Summarize PR**
```bash
git diffai pr 123
```

### 5. **Generate Changelog**
```bash
git diffai changelog
```

---

## 💡 Usage Examples

### **Commit Message Generation**
```bash
# Basic usage
git diffai commit

# Auto-commit with generated message
git diffai commit --auto

# Use detailed style
git diffai commit --style detailed

# Dry run (preview only)
git diffai commit --dry-run
```

### **PR Summarization**
```bash
# Basic PR summary
git diffai pr 123

# Include risk assessment
git diffai pr 123 --include-risk

# Output in markdown format
git diffai pr 123 --format markdown

# Verbose output with AI usage stats
git diffai pr 123 --verbose
```

### **Changelog Generation**
```bash
# Generate from last 10 commits
git diffai changelog

# Generate since specific tag
git diffai changelog --since v1.0.0

# Generate from last 20 commits
git diffai changelog --count 20

# Write to file
git diffai changelog --output CHANGELOG.md
```

### **Configuration Management**
```bash
# Show current configuration
git diffai config

# Initialize default config
git diffai config init

# Set specific values
git diffai config set ai.provider openai
git diffai config set ai.model gpt-4
```

---

## 🏗️ Architecture

DiffAI uses a **hybrid Go + Python architecture**:

```
┌─────────────────────────────────────────────────────────────┐
│                    DiffAI CLI (Go)                         │
├─────────────────────────────────────────────────────────────┤
│  ✅ Cobra CLI    ✅ Config     ✅ Git Engine               │
│  ✅ Commit Gen   ✅ PR Summary ✅ Changelog Gen            │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                AI Service Layer (Python)                   │
├─────────────────────────────────────────────────────────────┤
│  ✅ FastAPI      ✅ OpenAI     ✅ Anthropic                │
│  ✅ Local AI     ✅ Docker     ✅ Async Processing         │
└─────────────────────────────────────────────────────────────┘
```

**Why this architecture?**
- **Go CLI**: Fast execution, cross-platform, no external dependencies
- **Python AI Service**: Flexible AI integration, rapid development, rich ecosystem
- **Microservice**: Scalable, maintainable, language-agnostic

---

## 🛠️ Development

### **Prerequisites**
- Go 1.21+
- Python 3.11+
- Git
- Docker (optional)

### **Setup Development Environment**
```bash
# Clone repository
git clone https://github.com/obalaweb/diffai.git
cd diffai

# Setup development environment
make setup

# Start AI service for development
make dev

# Build CLI
make build

# Run tests
make test
```

### **Available Make Commands**
```bash
make build          # Build the CLI binary
make test           # Run tests
make dev            # Start development environment
make docker-build   # Build Docker images
make compose-up     # Start with Docker Compose
make lint           # Run linters
make fmt            # Format code
```

---

## 🛣️ Roadmap

### **v0.1.0 - MVP (Current)**
- [x] ✅ AI-powered commit message generation
- [x] ✅ PR summarization
- [x] ✅ Changelog generation
- [x] ✅ Multiple AI provider support
- [x] ✅ Configuration management
- [x] ✅ Docker support

### **v0.2.0 - Enhanced Features**
- [ ] 🔄 GitHub/GitLab integration
- [ ] 🔄 Pre-commit hooks
- [ ] 🔄 Custom templates
- [ ] 🔄 Batch processing
- [ ] 🔄 Team analytics

### **v0.3.0 - Advanced Features**
- [ ] 🔄 IDE extensions (VS Code, JetBrains)
- [ ] 🔄 Web interface
- [ ] 🔄 API rate limiting
- [ ] 🔄 Advanced diff analysis
- [ ] 🔄 Multi-language support

### **v1.0.0 - Production Ready**
- [ ] 🔄 Enterprise features
- [ ] 🔄 Advanced security
- [ ] 🔄 Performance optimization
- [ ] 🔄 Comprehensive testing
- [ ] 🔄 Production deployment guides

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### **Quick Contribution Guide**
1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### **Development Setup**
```bash
# Fork and clone
git clone https://github.com/your-username/diffai.git
cd diffai

# Setup development environment
make setup

# Start development
make dev
```

### **Areas for Contribution**
- 🐛 **Bug fixes** and improvements
- ✨ **New features** and enhancements
- 📚 **Documentation** improvements
- 🧪 **Tests** and test coverage
- 🎨 **UI/UX** improvements
- 🌍 **Internationalization**

---

## 📊 Project Status

| Component | Status | Progress |
|-----------|--------|----------|
| Architecture Design | ✅ Complete | 100% |
| Go CLI Framework | ✅ Complete | 100% |
| Git Integration | ✅ Complete | 100% |
| AI Service Layer | ✅ Complete | 100% |
| Configuration System | ✅ Complete | 100% |
| Development Environment | ✅ Complete | 100% |
| Testing Framework | 🚧 In Progress | 20% |
| CI/CD Pipeline | 🚧 In Progress | 10% |

---

## 🆘 Support

### **Getting Help**
- 📖 **Documentation**: Check our [docs](docs/) for detailed guides
- 🐛 **Bug Reports**: [Open an issue](https://github.com/obalaweb/diffai/issues)
- 💬 **Discussions**: [Join the conversation](https://github.com/obalaweb/diffai/discussions)
- 📧 **Contact**: [Create a discussion](https://github.com/obalaweb/diffai/discussions) for questions

### **Common Issues**
- **AI service not responding**: Check your API keys and network connection
- **Git integration issues**: Ensure you're in a Git repository
- **Permission errors**: Make sure DiffAI has proper file permissions

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- **OpenAI** for providing the GPT models
- **Anthropic** for Claude AI capabilities
- **go-git** for pure Go Git implementation
- **FastAPI** for the Python web framework
- **Cobra** for the Go CLI framework
- **All contributors** who help make DiffAI better

---

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=obalaweb/diffai&type=Date)](https://star-history.com/#obalaweb/diffai&Date)

---

<div align="center">

**Made with ❤️ by the DiffAI community**

[⭐ Star us on GitHub](https://github.com/obalaweb/diffai) • [🐛 Report Issues](https://github.com/obalaweb/diffai/issues) • [💬 Join Discussions](https://github.com/obalaweb/diffai/discussions) • [📖 Read Docs](docs/)

</div>
