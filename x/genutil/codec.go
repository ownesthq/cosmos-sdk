package genutil

import (
	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/auth"
	"github.com/ownesthq/cosmos-sdk/x/staking"
)

// generic sealed codec to be used throughout this module
var moduleCdc *codec.Codec

// TODO abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	moduleCdc = codec.New()
	staking.RegisterCodec(moduleCdc)
	auth.RegisterCodec(moduleCdc)
	sdk.RegisterCodec(moduleCdc)
	codec.RegisterCrypto(moduleCdc)
	moduleCdc.Seal()
}
