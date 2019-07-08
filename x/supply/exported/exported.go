package exported

import "github.com/ownesthq/cosmos-sdk/x/auth/exported"

// ModuleAccountI defines an account interface for modules that hold tokens in an escrow
type ModuleAccountI interface {
	exported.Account
	GetName() string
	GetPermission() string
}
