// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package starter provides a batteries-included bundle for omniagent.
//
// This package bundles commonly used skills and roles for quick agent setup:
//   - omniskill-pack: 18 markdown skills (GitHub CLI, weather, tmux, etc.)
//   - omniskill-github: GitHub SDK skill for issues, PRs, and code search
//   - omnirole-facilitator: Meeting facilitation role with workflows
//
// # Quick Start
//
// Use the default bundle for a fully-configured agent:
//
//	import starter "github.com/plexusone/omniagent-starter"
//
//	bundle := starter.Default(starter.Config{
//	    GitHubToken: os.Getenv("GITHUB_TOKEN"),
//	})
//
//	agent, _ := agent.New(config,
//	    agent.WithSkillPack(bundle.SkillPack()),
//	    agent.WithSkills(bundle.Skills()...),
//	    agent.WithRole(bundle.FacilitatorRole()),
//	)
//
// # Selective Usage
//
// Import only what you need:
//
//	bundle := starter.Default(starter.Config{})
//
//	// Just the skill pack (markdown skills)
//	agent.WithSkillPack(bundle.SkillPack())
//
//	// Just the GitHub skill
//	agent.WithSkills(bundle.GitHubSkill())
//
//	// Just the facilitator role
//	agent.WithRole(bundle.FacilitatorRole())
package starter

import (
	facilitator "github.com/plexusone/omnirole-facilitator"
	github "github.com/plexusone/omniskill-github"
	skills "github.com/plexusone/omniskill-pack"
	"github.com/plexusone/omniskill/pack"
	"github.com/plexusone/omniskill/role"
	"github.com/plexusone/omniskill/skill"
)

// Config configures the starter bundle.
type Config struct {
	// GitHub configuration
	GitHubToken        string
	GitHubBaseURL      string // For GitHub Enterprise
	GitHubDefaultOwner string
	GitHubDefaultRepo  string

	// Facilitator role configuration
	FacilitatorConfluenceSpace string
	FacilitatorAhaProduct      string
	FacilitatorTranscription   bool
	FacilitatorActionTracking  bool
}

// Bundle provides access to all bundled skills and roles.
type Bundle struct {
	config          Config
	skillPack       *skills.Pack
	githubSkill     *github.Skill
	facilitatorRole *facilitator.FacilitatorRole
}

// Default creates a new starter bundle with the given configuration.
func Default(cfg Config) *Bundle {
	b := &Bundle{
		config:    cfg,
		skillPack: skills.Default(),
	}

	// Create GitHub skill if token is provided
	if cfg.GitHubToken != "" {
		b.githubSkill = github.New(github.Config{
			Token:        cfg.GitHubToken,
			BaseURL:      cfg.GitHubBaseURL,
			DefaultOwner: cfg.GitHubDefaultOwner,
			DefaultRepo:  cfg.GitHubDefaultRepo,
		})
	}

	// Create facilitator role
	b.facilitatorRole = facilitator.New(facilitator.Config{
		DefaultConfluenceSpace: cfg.FacilitatorConfluenceSpace,
		DefaultAhaProduct:      cfg.FacilitatorAhaProduct,
		EnableTranscription:    cfg.FacilitatorTranscription,
		EnableActionTracking:   cfg.FacilitatorActionTracking,
	})

	return b
}

// SkillPack returns the omniskill-pack with 18 markdown skills.
func (b *Bundle) SkillPack() pack.SkillPack {
	return b.skillPack
}

// GitHubSkill returns the GitHub SDK skill.
// Returns nil if no GitHub token was configured.
func (b *Bundle) GitHubSkill() skill.Skill {
	if b.githubSkill == nil {
		return nil
	}
	return b.githubSkill
}

// FacilitatorRole returns the meeting facilitator role.
func (b *Bundle) FacilitatorRole() role.Role {
	return b.facilitatorRole
}

// Skills returns all configured skills as a slice.
// Excludes skills that weren't configured (e.g., GitHub without token).
func (b *Bundle) Skills() []skill.Skill {
	var result []skill.Skill
	if b.githubSkill != nil {
		result = append(result, b.githubSkill)
	}
	return result
}

// Roles returns all configured roles as a slice.
func (b *Bundle) Roles() []role.Role {
	return []role.Role{b.facilitatorRole}
}
