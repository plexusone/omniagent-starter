module github.com/plexusone/omniagent-starter

go 1.26.0

require (
	github.com/plexusone/omnirole-facilitator v0.1.0
	github.com/plexusone/omniskill v0.10.0
	github.com/plexusone/omniskill-github v0.1.0
	github.com/plexusone/omniskill-pack v0.1.0
)

require (
	github.com/google/go-github/v88 v88.0.0 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
)

// Local development - remove before tagging
replace (
	github.com/plexusone/omnirole-facilitator => ../omnirole-facilitator
	github.com/plexusone/omniskill-github => ../omniskill-github
	github.com/plexusone/omniskill-pack => ../omniskill-pack
)
