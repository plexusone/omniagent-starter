# OmniAgent Starter

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/omniagent-starter/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omniagent-starter
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omniagent-starter
 [viz-svg]: https://img.shields.io/badge/repo-visualization-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fomniagent-starter
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/omniagent-starter/blob/main/LICENSE

Batteries-included starter template for [OmniAgent](https://github.com/plexusone/omniagent). Run a fully-configured AI agent in minutes.

## Quick Start

```bash
# Clone the repo
git clone https://github.com/plexusone/omniagent-starter.git
cd omniagent-starter

# Set your API key
export OPENAI_API_KEY="sk-..."     # Or ANTHROPIC_API_KEY

# Enable WhatsApp
export WHATSAPP_ENABLED=true

# Run the agent
go run ./cmd/agent gateway run
```

A QR code will appear - scan it with WhatsApp to connect.

## What's Included

| Component | Description |
|-----------|-------------|
| **Runnable Agent** | `cmd/agent/main.go` - ready to run with standard omniagent CLI |
| **18 Markdown Skills** | Git, Docker, tmux, weather, and more via [omniskill-pack](https://github.com/plexusone/omniskill-pack) |
| **GitHub Skill** | Issues, PRs, code search via [omniskill-github](https://github.com/plexusone/omniskill-github) |
| **Web Search** | Google/news search via [omniserp](https://github.com/plexusone/omniserp) |
| **Facilitator Role** | Meeting notes and action tracking via [omnirole-facilitator](https://github.com/plexusone/omnirole-facilitator) |
| **Session Storage** | Persistent conversations with SQLite |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `OPENAI_API_KEY` | OpenAI API key |
| `ANTHROPIC_API_KEY` | Anthropic API key (alternative) |
| `WHATSAPP_ENABLED` | Set to `true` to enable WhatsApp |
| `GITHUB_TOKEN` | GitHub token for GitHub skill (optional) |
| `SERPER_API_KEY` | Serper API key for web search (optional) |
| `STORAGE_PATH` | SQLite path (default: `omniagent.db`) |
| `ENABLE_FACILITATOR_ROLE` | Set to `true` to use Meeting Facilitator role |

## Available Commands

After the agent starts, use standard omniagent commands:

```bash
# Gateway
go run ./cmd/agent gateway run      # Start with all channels

# Skills
go run ./cmd/agent skills list      # List bundled skills
go run ./cmd/agent skills check     # Check skill requirements

# Channels
go run ./cmd/agent channels status  # Show channel status

# Help
go run ./cmd/agent --help           # See all commands
```

## Customization

### Add Your Own Skills

Edit `cmd/agent/main.go` to add compiled skills:

```go
// Import your skill
import myskill "github.com/yourorg/myskill"

// Register it
mySkill := myskill.New(myskill.Config{...})
commands.RegisterAgentOption(agent.WithCompiledSkill(mySkill))
```

### Enable the Facilitator Role

```bash
export ENABLE_FACILITATOR_ROLE=true
go run ./cmd/agent gateway run
```

The agent will use the Meeting PM persona with workflows for meeting facilitation.

## Programmatic Usage

You can also import the starter bundle into your own agent:

```go
import (
    "github.com/plexusone/omniagent/agent"
    starter "github.com/plexusone/omniagent-starter"
)

bundle := starter.Default(starter.Config{
    GitHubToken: os.Getenv("GITHUB_TOKEN"),
})

agent, _ := agent.New(config,
    agent.WithSkillPack(bundle.SkillPack().FS()),
    agent.WithCompiledSkill(bundle.GitHubSkill()),
    agent.WithRole(bundle.FacilitatorRole()),
)
```

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  cmd/agent/main.go (runnable starter)                       │
├─────────────────────────────────────────────────────────────┤
│  starter.go (bundle)                                        │
│  ├── omniskill-pack         (18 markdown skills)            │
│  ├── omniskill-github       (GitHub SDK skill)              │
│  └── omnirole-facilitator   (meeting facilitation)          │
├─────────────────────────────────────────────────────────────┤
│  omniagent (core framework)                                 │
│  ├── Agent runtime with tool execution                      │
│  ├── WebSocket gateway                                      │
│  ├── Channel support (WhatsApp, Telegram, Discord, etc.)    │
│  └── Voice support (STT/TTS, realtime)                      │
└─────────────────────────────────────────────────────────────┘
```

## Requirements

- Go 1.26 or later

## Related

- [omniagent](https://github.com/plexusone/omniagent) - Core agent framework
- [grokify-omniagent](https://github.com/grokify/grokify-omniagent) - Example agent with investment skills

## License

MIT License - see [LICENSE](LICENSE) for details.
