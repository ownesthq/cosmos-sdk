package genutil

// DONTCOVER

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/auth"
	"github.com/ownesthq/cosmos-sdk/x/staking"
)

// ValidateAccountInGenesis checks that the provided key has sufficient
// coins in the genesis accounts
func ValidateAccountInGenesis(appGenesisState map[string]json.RawMessage,
	genAccIterator GenesisAccountsIterator,
	key sdk.AccAddress, coins sdk.Coins, cdc *codec.Codec) error {

	accountIsInGenesis := false

	// TODO refactor out bond denom to common state area
	stakingDataBz := appGenesisState[staking.ModuleName]
	var stakingData staking.GenesisState
	cdc.MustUnmarshalJSON(stakingDataBz, &stakingData)
	bondDenom := stakingData.Params.BondDenom

	genUtilDataBz := appGenesisState[staking.ModuleName]
	var genesisState GenesisState
	cdc.MustUnmarshalJSON(genUtilDataBz, &genesisState)

	var err error
	genAccIterator.IterateGenesisAccounts(cdc, appGenesisState,
		func(acc auth.Account) (stop bool) {
			accAddress := acc.GetAddress()
			accCoins := acc.GetCoins()

			// Ensure that account is in genesis
			if accAddress.Equals(key) {

				// Ensure account contains enough funds of default bond denom
				if coins.AmountOf(bondDenom).GT(accCoins.AmountOf(bondDenom)) {
					err = fmt.Errorf(
						"account %v is in genesis, but it only has %v%v available to stake, not %v%v",
						key.String(), accCoins.AmountOf(bondDenom), bondDenom, coins.AmountOf(bondDenom), bondDenom,
					)
					return true
				}
				accountIsInGenesis = true
				return true
			}
			return false
		},
	)
	if err != nil {
		return err
	}

	if !accountIsInGenesis {
		return fmt.Errorf("account %s in not in the app_state.accounts array of genesis.json", key)
	}

	return nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// deliver a genesis transaction
func DeliverGenTxs(ctx sdk.Context, cdc *codec.Codec, genTxs []json.RawMessage,
	stakingKeeper StakingKeeper, deliverTx deliverTxfn) []abci.ValidatorUpdate {

	for _, genTx := range genTxs {
		var tx auth.StdTx
		cdc.MustUnmarshalJSON(genTx, &tx)
		bz := cdc.MustMarshalBinaryLengthPrefixed(tx)
		res := deliverTx(abci.RequestDeliverTx{Tx: bz})
		if !res.IsOK() {
			panic(res.Log)
		}
	}
	return stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
}
