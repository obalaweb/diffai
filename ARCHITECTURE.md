# 🏗️ DiffAI Architecture Design

## 📋 Overview

DiffAI is designed as a **modular, extensible CLI tool** with a **hybrid Go + Python architecture** that leverages the strengths of both languages:

- **Go**: Fast CLI execution, Git integration, cross-platform compatibility
- **Python**: AI/ML processing, flexible model integration, rapid AI development

## 🎯 Design Principles

1. **Modularity**: Clear separation of concerns between CLI, Git operations, and AI processing
2. **Extensibility**: Plugin architecture for different AI providers and VCS platforms
3. **Performance**: Minimal latency for common operations, async processing for heavy tasks
4. **Reliability**: Robust error handling, graceful degradation, comprehensive testing
5. **Developer Experience**: Simple installation, intuitive CLI, helpful error messages

## 🏛️ High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    DiffAI CLI (Go)                         │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Cobra     │  │   Config    │  │   Git       │        │
│  │   CLI       │  │   Manager   │  │   Engine    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Diff      │  │   Commit    │  │   PR        │        │
│  │   Analyzer  │  │   Generator │  │   Summarizer│        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                AI Service Layer                             │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   HTTP      │  │   Python    │  │   AI        │        │
│  │   Client    │  │   Service   │  │   Providers │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│                              │                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   OpenAI    │  │ Anthropic   │  │   Local     │        │
│  │   API       │  │   Claude    │  │   Models    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
diffai/
├── cmd/
│   └── diffai/                 # Main CLI entry point
│       └── main.go
├── internal/
│   ├── cli/                    # CLI command implementations
│   │   ├── commit.go
│   │   ├── pr.go
│   │   ├── changelog.go
│   │   └── config.go
│   ├── git/                    # Git operations
│   │   ├── diff.go
│   │   ├── commit.go
│   │   └── repository.go
│   ├── ai/                     # AI service integration
│   │   ├── client.go
│   │   ├── providers/
│   │   │   ├── openai.go
│   │   │   ├── anthropic.go
│   │   │   └── local.go
│   │   └── models/
│   │       ├── commit.go
│   │       ├── pr.go
│   │       └── changelog.go
│   ├── config/                 # Configuration management
│   │   ├── config.go
│   │   └── defaults.go
│   └── utils/                  # Utility functions
│       ├── logger.go
│       ├── errors.go
│       └── validation.go
├── pkg/                        # Public API packages
│   ├── diffai/                 # Main package
│   └── types/                  # Common types
├── ai-service/                 # Python AI microservice
│   ├── app/
│   │   ├── __init__.py
│   │   ├── main.py
│   │   └── models/
│   │       ├── __init__.py
│   │       ├── commit.py
│   │       ├── pr.py
│   │       └── changelog.py
│   ├── services/
│   │   ├── __init__.py
│   │   ├── openai_service.py
│   │   ├── anthropic_service.py
│   │   └── local_service.py
│   ├── utils/
│   │   ├── __init__.py
│   │   ├── diff_parser.py
│   │   └── prompt_templates.py
│   ├── requirements.txt
│   └── Dockerfile
├── configs/                    # Configuration files
│   ├── default.yaml
│   └── templates/
├── scripts/                    # Build and deployment scripts
│   ├── build.sh
│   ├── test.sh
│   └── install.sh
├── tests/                      # Test files
│   ├── unit/
│   ├── integration/
│   └── fixtures/
├── docs/                       # Documentation
│   ├── api/
│   ├── user-guide/
│   └── development/
├── .github/                    # GitHub workflows
│   └── workflows/
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🔧 Core Components

### 1. CLI Layer (Go)

**Framework**: Cobra + Viper
- **Cobra**: Command structure, help system, argument parsing
- **Viper**: Configuration management, environment variables

**Commands**:
```bash
git diffai commit          # Generate commit message
git diffai pr <number>     # Summarize PR
git diffai changelog       # Generate changelog
git diffai config          # Manage configuration
git diffai version         # Show version info
```

### 2. Git Integration (Go)

**Library**: go-git
- Pure Go Git implementation
- No external Git dependency
- Cross-platform compatibility

**Key Operations**:
- Diff extraction and parsing
- Commit creation and management
- Repository state analysis
- Branch and PR information

### 3. AI Service Layer

**Architecture**: Microservice pattern
- **Communication**: HTTP REST API
- **Language**: Python (FastAPI)
- **Deployment**: Optional Docker container

**AI Providers**:
- OpenAI GPT models
- Anthropic Claude
- Local models (Ollama, etc.)
- Extensible plugin system

### 4. Configuration System

**Format**: YAML with environment variable overrides
**Location**: `~/.config/diffai/config.yaml`

**Configuration Schema**:
```yaml
ai:
  provider: "openai"  # openai, anthropic, local
  model: "gpt-4"
  api_key: "${OPENAI_API_KEY}"
  base_url: "https://api.openai.com/v1"
  
git:
  conventional_commits: true
  max_commit_length: 50
  auto_commit: false
  
output:
  format: "text"  # text, json, markdown
  verbose: false
```

## 🔄 Data Flow

### Commit Message Generation
```
1. User runs: git diffai commit
2. CLI extracts staged diff using go-git
3. CLI sends diff to AI service via HTTP
4. AI service processes diff with LLM
5. AI service returns structured commit message
6. CLI presents message to user for approval
7. CLI creates commit with generated message
```

### PR Summarization
```
1. User runs: git diffai pr 123
2. CLI fetches PR diff from Git provider
3. CLI sends diff + PR metadata to AI service
4. AI service generates comprehensive summary
5. CLI formats and displays summary
```

## 🚀 Deployment Options

### 1. Standalone Binary
- Single Go binary with embedded AI service
- No external dependencies
- Cross-platform distribution

### 2. Docker Container
- Go CLI + Python AI service
- Easy deployment and scaling
- Consistent environment

### 3. Cloud Service
- SaaS offering for teams
- API-based integration
- Advanced analytics and collaboration features

## 🔒 Security Considerations

1. **API Key Management**: Secure storage and transmission
2. **Code Privacy**: Local processing option for sensitive repos
3. **Rate Limiting**: Prevent API abuse
4. **Input Validation**: Sanitize all user inputs
5. **Audit Logging**: Track all AI service calls

## 📊 Performance Targets

- **Commit Generation**: < 2 seconds
- **PR Summarization**: < 5 seconds
- **Changelog Generation**: < 10 seconds
- **Memory Usage**: < 100MB for CLI
- **Binary Size**: < 50MB

## 🧪 Testing Strategy

### Unit Tests
- Go: Standard library + testify
- Python: pytest + pytest-asyncio

### Integration Tests
- Git repository fixtures
- AI service mocking
- End-to-end CLI workflows

### Performance Tests
- Load testing for AI service
- Memory profiling
- Response time benchmarks

## 🔮 Future Extensibility

### Plugin System
- Custom AI providers
- VCS integrations (GitLab, Bitbucket)
- Output formatters
- Pre-commit hooks

### IDE Integration
- VS Code extension
- JetBrains plugin
- Vim/Neovim support

### Team Features
- Shared configurations
- Commit style enforcement
- Collaboration analytics
- Custom templates

## 📈 Monitoring and Observability

- **Logging**: Structured logging with levels
- **Metrics**: Prometheus-compatible metrics
- **Tracing**: OpenTelemetry integration
- **Health Checks**: Service health endpoints

---

This architecture provides a solid foundation for building a professional, scalable, and maintainable AI-powered Git assistant while keeping the door open for future enhancements and community contributions.
