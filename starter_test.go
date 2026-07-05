// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package starter

import (
	"testing"
)

func TestDefault(t *testing.T) {
	bundle := Default(Config{})

	if bundle == nil {
		t.Fatal("Default() returned nil")
	}
}

func TestBundle_SkillPack(t *testing.T) {
	bundle := Default(Config{})

	pack := bundle.SkillPack()
	if pack == nil {
		t.Fatal("SkillPack() returned nil")
	}

	if pack.Name() != "omniskill-pack" {
		t.Errorf("SkillPack().Name() = %q, want %q", pack.Name(), "omniskill-pack")
	}
}

func TestBundle_GitHubSkill_NoToken(t *testing.T) {
	bundle := Default(Config{})

	skill := bundle.GitHubSkill()
	if skill != nil {
		t.Error("GitHubSkill() should be nil without token")
	}
}

func TestBundle_GitHubSkill_WithToken(t *testing.T) {
	bundle := Default(Config{
		GitHubToken: "test-token",
	})

	skill := bundle.GitHubSkill()
	if skill == nil {
		t.Fatal("GitHubSkill() returned nil with token")
	}

	if skill.Name() != "github" {
		t.Errorf("GitHubSkill().Name() = %q, want %q", skill.Name(), "github")
	}
}

func TestBundle_FacilitatorRole(t *testing.T) {
	bundle := Default(Config{})

	role := bundle.FacilitatorRole()
	if role == nil {
		t.Fatal("FacilitatorRole() returned nil")
	}

	if role.Name() != "facilitator" {
		t.Errorf("FacilitatorRole().Name() = %q, want %q", role.Name(), "facilitator")
	}
}

func TestBundle_Skills(t *testing.T) {
	// Without token - empty
	bundle := Default(Config{})
	skills := bundle.Skills()
	if len(skills) != 0 {
		t.Errorf("Skills() without token = %d skills, want 0", len(skills))
	}

	// With token - includes GitHub
	bundle = Default(Config{GitHubToken: "test-token"})
	skills = bundle.Skills()
	if len(skills) != 1 {
		t.Errorf("Skills() with token = %d skills, want 1", len(skills))
	}
}

func TestBundle_Roles(t *testing.T) {
	bundle := Default(Config{})

	roles := bundle.Roles()
	if len(roles) != 1 {
		t.Errorf("Roles() = %d roles, want 1", len(roles))
	}
}
