package keeper

import (
	sdk "github.com/ownesthq/cosmos-sdk/types"
)

// CrisisKeeper defines the expected crisis keeper
type CrisisKeeper interface {
	RegisterRoute(moduleName, route string, invar sdk.Invariant)
}
