// nolint
package cli

import (
	"github.com/spf13/cobra"

	"github.com/ownesthq/cosmos-sdk/client"
	"github.com/ownesthq/cosmos-sdk/client/context"
	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/auth"
	"github.com/ownesthq/cosmos-sdk/x/auth/client/utils"
	"github.com/ownesthq/cosmos-sdk/x/crisis/internal/types"
)

// command to replace a delegator's withdrawal address
func GetCmdInvariantBroken(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invariant-broken [module-name] [invariant-route]",
		Short: "submit proof that an invariant broken to halt the chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			senderAddr := cliCtx.GetFromAddress()
			moduleName, route := args[0], args[1]
			msg := types.NewMsgVerifyInvariant(senderAddr, moduleName, route)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Crisis transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(client.PostCommands(
		GetCmdInvariantBroken(cdc),
	)...)
	return txCmd
}
