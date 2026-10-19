package main

import (
	"github.com/brickKit/be-acceptance/probes/go/v3/probe"
	"github.com/brickKit/be-sdk-go/v3/component"
	audit "github.com/brickKit/infra-audit/v3/member"
	authz "github.com/brickKit/infra-authz/v3/member"
	"github.com/brickKit/infra-iam-casdoor/v3/iam"
)

// members are the definitions compiled into this shell, at the versions go.mod pins (by tag, never replace); they
// must equal shell.members in component.yaml (TestMembersMatchTheManifest).
func members() []component.Definition {
	// the experiment's probes and step 5's security members; business members join in steps 6-7
	return append(probe.Definitions(), iam.Definition(), authz.Definition(), audit.Definition())
}
