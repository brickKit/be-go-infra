package main

import (
	"github.com/brickKit/be-acceptance/probes/go/v3/probe"
	"github.com/brickKit/be-sdk-go/v3/component"
)

// members are the definitions compiled into this shell, at the versions go.mod pins (by tag, never replace); they
// must equal shell.members in component.yaml (TestMembersMatchTheManifest).
func members() []component.Definition {
	return probe.Definitions() // the experiment's probes; business members join in steps 5-7
}
