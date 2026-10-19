package main

import (
	"github.com/brickKit/be-acceptance/probes/go/v3/probe"
	"github.com/brickKit/be-sdk-go/v3/component"
	audit "github.com/brickKit/infra-audit/v3/member"
	authz "github.com/brickKit/infra-authz/v3/member"
	"github.com/brickKit/infra-iam-casdoor/v3/iam"
	"github.com/brickKit/mdm-currency/v3/currency"
	"github.com/brickKit/mdm-org/v3/org"
	"github.com/brickKit/mdm-partner/v3/partner"
	"github.com/brickKit/mdm-product/v3/product"
)

// members are the definitions compiled into this shell, at the versions go.mod pins (by tag, never replace); they
// must equal shell.members in component.yaml (TestMembersMatchTheManifest).
func members() []component.Definition {
	// the experiment's probes, step 5's security members and step 6's master data; the business members of step 7
	// go to be-go-core
	return append(probe.Definitions(), iam.Definition(), authz.Definition(), audit.Definition(), org.Definition(),
		currency.Definition(), partner.Definition(), product.Definition())
}
