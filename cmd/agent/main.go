// Package main provides a starter agent demonstrating omniagent features.
//
// This is a batteries-included agent that shows how to configure:
//   - Skill packs (bundled markdown skills)
//   - Compiled skills (GitHub, web search)
//   - Roles (Meeting Facilitator)
//   - Persistent sessions
//
// # Quick Start
//
// Set your API keys and run:
//
//	export OPENAI_API_KEY="sk-..."           # Or ANTHROPIC_API_KEY
//	export GITHUB_TOKEN="ghp_..."            # Optional: for GitHub skill
//	export SERPER_API_KEY="..."              # Optional: for web search
//	export WHATSAPP_ENABLED=true             # Enable WhatsApp
//	go run ./cmd/agent gateway run
//
// # Environment Variables
//
//	OMNIAGENT_AGENT_PROVIDER  - LLM provider: openai, anthropic, gemini (default: openai)
//	OMNIAGENT_AGENT_MODEL     - Model name (default: gpt-4o)
//	STORAGE_PATH              - SQLite path for sessions (default: omniagent.db)
//	GITHUB_TOKEN              - GitHub token for GitHub skill
//	SERPER_API_KEY            - Serper API key for web search
//	ENABLE_FACILITATOR_ROLE   - Set to "true" to use Meeting Facilitator role
//
// # Available Commands
//
// After building, use standard omniagent commands:
//
//	./agent gateway run      # Start the gateway with all channels
//	./agent skills list      # List bundled skills
//	./agent channels status  # Show channel status
//	./agent version          # Show version info
package main

import (
	"log"
	"os"

	"github.com/plexusone/omniagent/agent"
	"github.com/plexusone/omniagent/cmd/omniagent/commands"
	"github.com/plexusone/omniagent/skills/compiled"
	"github.com/plexusone/omnistorage-core/kvs/backend/sqlite"

	// Starter bundle with skills and roles
	starter "github.com/plexusone/omniagent-starter"

	// Additional compiled skills
	searchskill "github.com/plexusone/omniserp/omniskill"
)

func main() {
	// =========================================================================
	// 1. Storage Setup (for persistent sessions)
	// =========================================================================
	storagePath := getEnv("STORAGE_PATH", "omniagent.db")
	store, err := sqlite.New(sqlite.Config{
		Path: storagePath,
	})
	if err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}
	commands.RegisterAgentOption(agent.WithSessionsFromStorage(store))
	log.Printf("Sessions stored in: %s", storagePath)

	// =========================================================================
	// 2. Create Starter Bundle
	// =========================================================================
	// The starter bundle provides:
	//   - SkillPack(): 18 markdown skills (git, docker, tmux, weather, etc.)
	//   - GitHubSkill(): GitHub SDK skill for issues, PRs, code search
	//   - FacilitatorRole(): Meeting facilitation with workflows
	bundle := starter.Default(starter.Config{
		// GitHub configuration (optional)
		GitHubToken: os.Getenv("GITHUB_TOKEN"),

		// Facilitator role configuration
		FacilitatorActionTracking: true,
	})

	// =========================================================================
	// 3. Register Skill Pack (Markdown Skills)
	// =========================================================================
	// The skill pack includes 18 commonly-used skills in SKILL.md format.
	// These are injected into the system prompt automatically.
	commands.RegisterAgentOption(agent.WithSkillPack(bundle.SkillPack().FS()))
	log.Println("Registered skill pack with 18 markdown skills")

	// =========================================================================
	// 4. Register Compiled Skills
	// =========================================================================

	// GitHub skill (if token provided)
	// Cast to compiled.Skill since skill.Skill is interface-compatible
	if ghSkill := bundle.GitHubSkill(); ghSkill != nil {
		commands.RegisterAgentOption(agent.WithCompiledSkill(ghSkill.(compiled.Skill)))
		log.Println("Registered GitHub skill")
	}

	// Web search skill (if API key provided)
	if os.Getenv("SERPER_API_KEY") != "" || os.Getenv("SERPAPI_API_KEY") != "" {
		searchSkill, err := searchskill.New(searchskill.Config{})
		if err != nil {
			log.Printf("Warning: failed to create search skill: %v", err)
		} else {
			commands.RegisterAgentOption(agent.WithCompiledSkill(searchSkill))
			log.Println("Registered web search skill")
		}
	}

	// =========================================================================
	// 5. Register Role (Optional)
	// =========================================================================
	// Roles provide high-level personas that combine:
	//   - System prompt customization
	//   - Workflows (structured multi-step operations)
	//   - Policies (tool access control)
	//   - Behaviors (context-aware actions)
	if os.Getenv("ENABLE_FACILITATOR_ROLE") == "true" {
		commands.RegisterAgentOption(agent.WithRole(bundle.FacilitatorRole()))
		log.Println("Using Meeting Facilitator role")
	}

	// =========================================================================
	// 6. Set Defaults
	// =========================================================================

	// Default voice greeting
	if os.Getenv("OMNIAGENT_VOICE_GREETING") == "" {
		_ = os.Setenv("OMNIAGENT_VOICE_GREETING",
			"Hello! I'm your OmniAgent assistant. How can I help you today?")
	}

	// Default provider if not set
	if os.Getenv("OMNIAGENT_AGENT_PROVIDER") == "" {
		_ = os.Setenv("OMNIAGENT_AGENT_PROVIDER", "openai")
	}

	// Default model if not set
	if os.Getenv("OMNIAGENT_AGENT_MODEL") == "" {
		_ = os.Setenv("OMNIAGENT_AGENT_MODEL", "gpt-4o")
	}

	// =========================================================================
	// 7. Run OmniAgent CLI
	// =========================================================================
	// This runs the standard omniagent CLI with all registered skills/roles.
	// Configuration comes from config file or environment variables.
	log.Println("Starting OmniAgent...")
	if err := commands.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
