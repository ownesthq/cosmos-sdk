package gov

import (
	"fmt"

	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/gov/types"
)

// RegisterInvariants registers all governance invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "module-account", ModuleAccountInvariant(keeper))
}

// AllInvariants runs all invariants of the governance module
func AllInvariants(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) error {
		return ModuleAccountInvariant(keeper)(ctx)
	}
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// deposit amounts held on store
func ModuleAccountInvariant(keeper Keeper) sdk.Invariant {
	return func(ctx sdk.Context) error {
		var expectedDeposits sdk.Coins

		keeper.IterateAllDeposits(ctx, func(deposit types.Deposit) bool {
			expectedDeposits = expectedDeposits.Add(deposit.Amount)
			return false
		})

		macc := keeper.GetGovernanceAccount(ctx)
		if !macc.GetCoins().IsEqual(expectedDeposits) {
			return fmt.Errorf("deposits invariance:\n"+
				"\tgov ModuleAccount coins: %s\n"+
				"\tsum of deposit amounts: %s", macc.GetCoins(), expectedDeposits)
		}

		return nil
	}
}
