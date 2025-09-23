# Changelog

All notable changes to DiffAI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- N/A

### Changed
- N/A

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

### Security
- N/A

## [0.1.0] - 2025-09-23

### Added
- Initial release of DiffAI - AI-powered Git assistant
- AI-powered commit message generation from staged diffs
- Pull request summarization with impact analysis
- Automated changelog generation from commit history
- Multiple AI provider support (OpenAI GPT-4, Anthropic Claude, Local AI)
- Git CLI integration using go-git (pure Go implementation)
- Comprehensive configuration management system
- Docker support and development environment
- Cross-platform compatibility (Linux, macOS, Windows)
- Hybrid Go + Python microservice architecture
- FastAPI-based AI service layer
- Cobra CLI framework with all core commands
- Comprehensive documentation and contributing guidelines
- Professional README with installation and usage examples
- Development tooling with Makefile and Docker Compose

---

## Release Notes

### v0.1.0 - Initial Release

This is the initial release of DiffAI, an AI-powered Git assistant that helps developers create better commits, generate meaningful PR summaries, and automate changelogs.

**Key Features:**
- 🤖 AI-powered commit message generation from staged diffs
- 📑 Comprehensive PR summaries with risk assessment
- 📝 Automated changelog generation from commit history
- 🔌 Seamless Git CLI integration (`git diffai commit`)
- 🌍 Support for multiple AI providers (OpenAI, Anthropic, Local)
- ⚙️ Flexible configuration system
- 🐳 Docker support for easy deployment
- 📚 Comprehensive documentation and examples

**Installation:**
```bash
# Download binary from releases
curl -L https://github.com/diffai/diffai/releases/latest/download/diffai-linux-amd64 -o diffai
chmod +x diffai
sudo mv diffai /usr/local/bin/

# Or build from source
git clone https://github.com/diffai/diffai.git
cd diffai
make build
sudo make install
```

**Quick Start:**
```bash
# Initialize configuration
git diffai config init

# Generate commit message
git add .
git diffai commit

# Summarize PR
git diffai pr 123

# Generate changelog
git diffai changelog
```

**Configuration:**
Set your AI provider API key:
```bash
export OPENAI_API_KEY="your-api-key"
# or
export ANTHROPIC_API_KEY="your-api-key"
```

**Documentation:**
- [User Guide](docs/user-guide/)
- [API Reference](docs/api/)
- [Development Guide](docs/development/)
- [Contributing Guidelines](CONTRIBUTING.md)

**Support:**
- GitHub Issues: [Report bugs and request features](https://github.com/diffai/diffai/issues)
- GitHub Discussions: [Ask questions and share ideas](https://github.com/diffai/diffai/discussions)
- Documentation: [Read the docs](https://github.com/diffai/diffai/tree/main/docs)

---

## Version History

| Version | Date | Description |
|---------|------|-------------|
| 0.1.0 | 2025-01-XX | Initial release with core functionality |

## Migration Guide

### From Pre-release to v0.1.0

This is the first stable release, so no migration is needed. If you were using a pre-release version:

1. Update to the latest release
2. Run `git diffai config init` to create a new configuration file
3. Set your AI provider API key in the environment or configuration file

---

*For more information about DiffAI, visit [https://github.com/diffai/diffai](https://github.com/diffai/diffai)*
