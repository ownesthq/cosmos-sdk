// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/ownesthq/cosmos-sdk/x/auth/types
package auth

import (
	"github.com/ownesthq/cosmos-sdk/x/auth/types"
)

const (
	ModuleName                    = types.ModuleName
	StoreKey                      = types.StoreKey
	FeeStoreKey                   = types.FeeStoreKey
	QuerierRoute                  = types.QuerierRoute
	DefaultParamspace             = types.DefaultParamspace
	DefaultMaxMemoCharacters      = types.DefaultMaxMemoCharacters
	DefaultTxSigLimit             = types.DefaultTxSigLimit
	DefaultTxSizeCostPerByte      = types.DefaultTxSizeCostPerByte
	DefaultSigVerifyCostED25519   = types.DefaultSigVerifyCostED25519
	DefaultSigVerifyCostSecp256k1 = types.DefaultSigVerifyCostSecp256k1
	QueryAccount                  = types.QueryAccount
)

var (
	// functions aliases
	NewBaseAccount                 = types.NewBaseAccount
	ProtoBaseAccount               = types.ProtoBaseAccount
	NewBaseAccountWithAddress      = types.NewBaseAccountWithAddress
	NewBaseVestingAccount          = types.NewBaseVestingAccount
	NewContinuousVestingAccountRaw = types.NewContinuousVestingAccountRaw
	NewContinuousVestingAccount    = types.NewContinuousVestingAccount
	NewDelayedVestingAccountRaw    = types.NewDelayedVestingAccountRaw
	NewDelayedVestingAccount       = types.NewDelayedVestingAccount
	RegisterCodec                  = types.RegisterCodec
	RegisterBaseAccount            = types.RegisterBaseAccount
	NewFeeCollectionKeeper         = types.NewFeeCollectionKeeper
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	ValidateGenesis                = types.ValidateGenesis
	AddressStoreKey                = types.AddressStoreKey
	NewParams                      = types.NewParams
	ParamKeyTable                  = types.ParamKeyTable
	DefaultParams                  = types.DefaultParams
	NewQueryAccountParams          = types.NewQueryAccountParams
	NewStdTx                       = types.NewStdTx
	CountSubKeys                   = types.CountSubKeys
	NewStdFee                      = types.NewStdFee
	StdSignBytes                   = types.StdSignBytes
	DefaultTxDecoder               = types.DefaultTxDecoder
	DefaultTxEncoder               = types.DefaultTxEncoder
	NewTxBuilder                   = types.NewTxBuilder
	NewTxBuilderFromCLI            = types.NewTxBuilderFromCLI
	MakeSignature                  = types.MakeSignature

	// variable aliases
	ModuleCdc                 = types.ModuleCdc
	AddressStoreKeyPrefix     = types.AddressStoreKeyPrefix
	GlobalAccountNumberKey    = types.GlobalAccountNumberKey
	KeyMaxMemoCharacters      = types.KeyMaxMemoCharacters
	KeyTxSigLimit             = types.KeyTxSigLimit
	KeyTxSizeCostPerByte      = types.KeyTxSizeCostPerByte
	KeySigVerifyCostED25519   = types.KeySigVerifyCostED25519
	KeySigVerifyCostSecp256k1 = types.KeySigVerifyCostSecp256k1
)

type (
	Account                  = types.Account
	VestingAccount           = types.VestingAccount
	AccountDecoder           = types.AccountDecoder
	BaseAccount              = types.BaseAccount
	BaseVestingAccount       = types.BaseVestingAccount
	ContinuousVestingAccount = types.ContinuousVestingAccount
	DelayedVestingAccount    = types.DelayedVestingAccount
	FeeCollectionKeeper      = types.FeeCollectionKeeper
	GenesisState             = types.GenesisState
	Params                   = types.Params
	QueryAccountParams       = types.QueryAccountParams
	StdSignMsg               = types.StdSignMsg
	StdTx                    = types.StdTx
	StdFee                   = types.StdFee
	StdSignDoc               = types.StdSignDoc
	StdSignature             = types.StdSignature
	TxBuilder                = types.TxBuilder
)
