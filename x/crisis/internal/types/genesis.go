package types

import (
	"fmt"

	sdk "github.com/ownesthq/cosmos-sdk/types"
)

// GenesisState - crisis genesis state
type GenesisState struct {
	ConstantFee sdk.Coin `json:"constant_fee"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(constantFee sdk.Coin) GenesisState {
	return GenesisState{
		ConstantFee: constantFee,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		ConstantFee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)),
	}
}

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(data GenesisState) error {
	if !data.ConstantFee.IsPositive() {
		return fmt.Errorf("constant fee must be positive: %s", data.ConstantFee)
	}
	return nil
}
