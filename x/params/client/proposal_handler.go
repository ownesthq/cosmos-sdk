package client

import (
	govclient "github.com/ownesthq/cosmos-sdk/x/gov/client"
	"github.com/ownesthq/cosmos-sdk/x/params/client/cli"
	"github.com/ownesthq/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
