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

A batteries-included starter template for building AI agents with [OmniAgent](https://github.com/plexusone/omniagent). Get a fully-configured agent running in minutes, then customize it for your needs.

## What is OmniAgent?

**OmniAgent** is a Go framework for building AI agents that can communicate across multiple channels (WhatsApp, Telegram, Discord, phone calls, WebRTC meetings) and perform tasks using tools. Think of it as your AI representative that can:

- Respond to messages on your behalf across messaging platforms
- Execute tools like web search, GitHub operations, or custom integrations
- Have voice conversations via phone (Twilio/Telnyx) or browser (LiveKit WebRTC)
- Remember context across conversations with persistent sessions
- Take on specialized personas with roles and workflows

This starter template gives you a working agent out of the box, pre-configured with commonly used skills and ready to connect to WhatsApp.

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

A QR code will appear in your terminal. Open WhatsApp on your phone, go to **Settings → Linked Devices → Link a Device**, and scan the QR code. Your agent is now connected and will respond to messages.

## Understanding Skills and Roles

OmniAgent uses two key concepts for extending agent capabilities:

### Skills

**Skills** give your agent specific capabilities. There are two types:

- **Markdown Skills** - Instructions in `SKILL.md` files that get injected into the system prompt. They tell the LLM how to use external tools (like `gh` CLI for GitHub or `docker` for containers).

- **Compiled Skills** - Go code that registers tools directly with the LLM. These provide type-safe function calling with proper parameter validation.

This starter includes both:

| Skill | Type | What it does |
|-------|------|--------------|
| 18 skills from [omniskill-pack](https://github.com/plexusone/omniskill-pack) | Markdown | Git, Docker, tmux, weather, Homebrew, and more |
| [omni-github](https://github.com/plexusone/omni-github) | Compiled | Search issues, PRs, and code via GitHub API |
| [omniserp](https://github.com/plexusone/omniserp) | Compiled | Web and news search via Serper/SerpAPI |

### Roles

**Roles** are higher-level personas that combine skills with specialized behavior. A role defines:

- A system prompt that shapes the agent's personality and expertise
- Workflows for structured multi-step operations
- Policies for tool access control
- Context-aware behaviors (different actions in meetings vs chat)

This starter includes the **Meeting Facilitator** role from [omnirole-facilitator](https://github.com/plexusone/omnirole-facilitator), which turns your agent into a meeting assistant that can take notes, track action items, and manage follow-ups.

## What's Included

| Component | Description |
|-----------|-------------|
| `cmd/agent/main.go` | Runnable agent with all skills pre-registered |
| `starter.go` | Bundle for programmatic use in your own agents |
| 18 Markdown Skills | Git, Docker, tmux, Homebrew, weather, and more |
| GitHub Skill | Issues, PRs, code search via GitHub API |
| Web Search Skill | Google/news search via Serper |
| Facilitator Role | Meeting notes and action tracking |
| Session Storage | Persistent conversations with SQLite |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `OPENAI_API_KEY` | OpenAI API key for GPT models |
| `ANTHROPIC_API_KEY` | Anthropic API key for Claude models (alternative) |
| `GEMINI_API_KEY` | Google API key for Gemini models (alternative) |
| `WHATSAPP_ENABLED` | Set to `true` to enable WhatsApp channel |
| `TELEGRAM_BOT_TOKEN` | Telegram bot token (auto-enables Telegram) |
| `DISCORD_BOT_TOKEN` | Discord bot token (auto-enables Discord) |
| `GITHUB_TOKEN` | GitHub token for GitHub skill (optional) |
| `SERPER_API_KEY` | Serper API key for web search (optional) |
| `STORAGE_PATH` | SQLite path for sessions (default: `omniagent.db`) |
| `ENABLE_FACILITATOR_ROLE` | Set to `true` to use Meeting Facilitator role |

## Available Commands

The agent uses the standard OmniAgent CLI. After building, you have access to:

```bash
# Start the gateway (connects to messaging channels)
go run ./cmd/agent gateway run

# List registered skills
go run ./cmd/agent skills list

# Check which skills have missing requirements
go run ./cmd/agent skills check

# Show channel connection status
go run ./cmd/agent channels status

# Interactive setup wizard
go run ./cmd/agent setup

# Diagnose configuration issues
go run ./cmd/agent doctor

# Start voice gateway for phone calls
go run ./cmd/agent voice serve --provider twilio

# See all commands
go run ./cmd/agent --help
```

## Customization

### Adding Your Own Skills

Edit `cmd/agent/main.go` to register additional skills:

```go
// Import your skill package
import myskill "github.com/yourorg/myskill"

// In main(), register it with the agent
mySkill := myskill.New(myskill.Config{
    APIKey: os.Getenv("MYSKILL_API_KEY"),
})
commands.RegisterAgentOption(agent.WithCompiledSkill(mySkill))
```

Skills implement a simple interface:

```go
type Skill interface {
    Name() string
    Description() string
    Tools() []Tool
    Init(ctx context.Context) error
    Close() error
}
```

### Enabling the Facilitator Role

To use the Meeting Facilitator persona:

```bash
export ENABLE_FACILITATOR_ROLE=true
go run ./cmd/agent gateway run
```

The agent will introduce itself as a meeting facilitator and offer to help with agendas, note-taking, and action item tracking.

### Connecting Multiple Channels

Enable additional channels with environment variables:

```bash
# WhatsApp (QR code linking)
export WHATSAPP_ENABLED=true

# Telegram (get token from @BotFather)
export TELEGRAM_BOT_TOKEN="123456:ABC-DEF..."

# Discord (create bot at discord.com/developers)
export DISCORD_BOT_TOKEN="MTIz..."

go run ./cmd/agent gateway run
```

## Programmatic Usage

You can import the starter bundle into your own agent instead of using the CLI:

```go
import (
    "os"

    "github.com/plexusone/omniagent/agent"
    starter "github.com/plexusone/omniagent-starter"
)

func main() {
    // Create the starter bundle
    bundle := starter.Default(starter.Config{
        GitHubToken:               os.Getenv("GITHUB_TOKEN"),
        FacilitatorActionTracking: true,
    })

    // Create agent with bundled skills
    a, _ := agent.New(config,
        agent.WithSkillPack(bundle.SkillPack().FS()),
        agent.WithCompiledSkill(bundle.GitHubSkill()),
        agent.WithRole(bundle.FacilitatorRole()),
    )

    // Use the agent...
}
```

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  Your Messages (WhatsApp, Telegram, Discord, Phone, etc.)   │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  cmd/agent/main.go (this starter)                           │
│  - Registers skills and roles                               │
│  - Configures sessions and storage                          │
│  - Runs the OmniAgent CLI                                   │
├─────────────────────────────────────────────────────────────┤
│  starter.go (bundle)                                        │
│  ├── omniskill-pack         (18 markdown skills)            │
│  ├── omni-github            (GitHub API skill)              │
│  ├── omniserp               (web search skill)              │
│  └── omnirole-facilitator   (meeting facilitator role)      │
├─────────────────────────────────────────────────────────────┤
│  omniagent (core framework)                                 │
│  ├── Agent runtime - LLM integration, tool execution        │
│  ├── Gateway - WebSocket control plane, REST API            │
│  ├── Channels - WhatsApp, Telegram, Discord, Twilio SMS     │
│  ├── Voice - Phone calls (Twilio/Telnyx), WebRTC (LiveKit)  │
│  └── Sessions - Conversation history, SQLite storage        │
├─────────────────────────────────────────────────────────────┤
│  omnillm (LLM providers)                                    │
│  └── OpenAI, Anthropic, Google Gemini, AWS Bedrock, etc.    │
└─────────────────────────────────────────────────────────────┘
```

## The Omni* Ecosystem

OmniAgent is built on a modular ecosystem of `omni*` packages:

| Package | Purpose |
|---------|---------|
| [omniagent](https://github.com/plexusone/omniagent) | Core agent framework |
| [omnillm](https://github.com/plexusone/omnillm) | Multi-provider LLM abstraction |
| [omniskill](https://github.com/plexusone/omniskill) | Skill interfaces and types |
| [omnichat](https://github.com/plexusone/omnichat) | Messaging channel abstraction |
| [omnivoice](https://github.com/plexusone/omnivoice) | Voice STT/TTS interfaces |
| [omnimemory](https://github.com/plexusone/omnimemory) | Semantic memory with vector search |

## Requirements

- Go 1.26 or later
- An LLM API key (OpenAI, Anthropic, or Google)

## License

MIT License - see [LICENSE](LICENSE) for details.
