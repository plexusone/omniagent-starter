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

Batteries-included bundle for OmniAgent with commonly used skills and roles.

## What's Included

| Package | Description |
|---------|-------------|
| [omniskill-pack](https://github.com/plexusone/omniskill-pack) | 18 markdown skills (GitHub CLI, weather, tmux, etc.) |
| [omniskill-github](https://github.com/plexusone/omniskill-github) | GitHub SDK skill for issues, PRs, code search |
| [omnirole-facilitator](https://github.com/plexusone/omnirole-facilitator) | Meeting facilitation role with workflows |

## Installation

```bash
go get github.com/plexusone/omniagent-starter
```

## Quick Start

```go
import (
    "os"

    "github.com/plexusone/omniagent/agent"
    starter "github.com/plexusone/omniagent-starter"
)

func main() {
    // Create bundle with configuration
    bundle := starter.Default(starter.Config{
        GitHubToken: os.Getenv("GITHUB_TOKEN"),
    })

    // Use with omniagent
    agent, _ := agent.New(config,
        agent.WithSkillPack(bundle.SkillPack()),
        agent.WithSkills(bundle.Skills()...),
        agent.WithRole(bundle.FacilitatorRole()),
    )
}
```

## Selective Usage

Import only what you need:

```go
bundle := starter.Default(starter.Config{})

// Just the skill pack (18 markdown skills)
agent.WithSkillPack(bundle.SkillPack())

// Just the GitHub skill
agent.WithSkills(bundle.GitHubSkill())

// Just the facilitator role
agent.WithRole(bundle.FacilitatorRole())
```

## Configuration

```go
starter.Config{
    // GitHub skill configuration
    GitHubToken:        "ghp_...",           // Required for GitHub skill
    GitHubBaseURL:      "https://...",       // For GitHub Enterprise
    GitHubDefaultOwner: "myorg",             // Default repo owner
    GitHubDefaultRepo:  "myrepo",            // Default repo name

    // Facilitator role configuration
    FacilitatorConfluenceSpace: "TEAM",      // Confluence space for notes
    FacilitatorAhaProduct:      "PROD-1",    // Aha product ID
    FacilitatorTranscription:   true,        // Enable transcription
    FacilitatorActionTracking:  true,        // Track action items
}
```

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  Your Agent (custom configuration)                          │
├─────────────────────────────────────────────────────────────┤
│  omniagent-starter (this package)                           │
│  ├── omniskill-pack         (18 markdown skills)            │
│  ├── omniskill-github       (GitHub SDK skill)              │
│  └── omnirole-facilitator   (meeting facilitation)          │
├─────────────────────────────────────────────────────────────┤
│  omniagent (core framework)                                 │
├─────────────────────────────────────────────────────────────┤
│  omniskill (interfaces & types)                             │
└─────────────────────────────────────────────────────────────┘
```

## Requirements

- Go 1.26 or later

## License

MIT License - see [LICENSE](LICENSE) for details.
