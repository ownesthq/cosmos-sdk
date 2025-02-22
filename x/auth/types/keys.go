package types

import sdk "github.com/ownesthq/cosmos-sdk/types"

const (
	// module name
	ModuleName = "auth"

	// StoreKey is string representation of the store key for auth
	StoreKey = "acc"

	// FeeStoreKey is a string representation of the store key for fees
	FeeStoreKey = "fee"

	// QuerierRoute is the querier route for acc
	QuerierRoute = StoreKey
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	AddressStoreKeyPrefix = []byte{0x01}

	// param key for global account number
	GlobalAccountNumberKey = []byte("globalAccountNumber")
)

// AddressStoreKey turn an address to key used to get it from the account store
func AddressStoreKey(addr sdk.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}
