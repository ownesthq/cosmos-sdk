package genaccounts

import (
	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
)

// initialize accounts and deliver genesis transactions
func InitGenesis(ctx sdk.Context, _ *codec.Codec, accountKeeper AccountKeeper, genesisState GenesisState) {
	genesisState.Sanitize()

	// load the accounts
	for _, gacc := range genesisState {
		acc := gacc.ToAccount()
		acc = accountKeeper.NewAccount(ctx, acc) // set account number
		accountKeeper.SetAccount(ctx, acc)
	}
}
