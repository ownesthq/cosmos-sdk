package genaccounts

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ownesthq/cosmos-sdk/client/context"
	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/types/module"
	"github.com/ownesthq/cosmos-sdk/x/auth"
)

var (
	_ module.AppModuleGenesis = AppModule{}
	_ module.AppModuleBasic   = AppModuleBasic{}
)

// module name
const ModuleName = "accounts"

// app module basics object
type AppModuleBasic struct{}

// module name
func (AppModuleBasic) Name() string {
	return ModuleName
}

// register module codec
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {}

// default genesis state
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return moduleCdc.MustMarshalJSON(GenesisState{})
}

// module validate genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := moduleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// register rest routes
func (AppModuleBasic) RegisterRESTRoutes(_ context.CLIContext, _ *mux.Router) {}

// get the root tx command of this module
func (AppModuleBasic) GetTxCmd(_ *codec.Codec) *cobra.Command { return nil }

// get the root query command of this module
func (AppModuleBasic) GetQueryCmd(_ *codec.Codec) *cobra.Command { return nil }

// extra function from sdk.AppModuleBasic
// iterate the genesis accounts and perform an operation at each of them
// - to used by other modules
func (AppModuleBasic) IterateGenesisAccounts(cdc *codec.Codec, appGenesis map[string]json.RawMessage,
	iterateFn func(auth.Account) (stop bool)) {

	genesisState := GetGenesisStateFromAppState(cdc, appGenesis)
	for _, genAcc := range genesisState {
		acc := genAcc.ToAccount()
		if iterateFn(acc) {
			break
		}
	}
}

//___________________________
// app module
type AppModule struct {
	AppModuleBasic
	accountKeeper AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(accountKeeper AccountKeeper) module.AppModule {

	return module.NewGenesisOnlyAppModule(AppModule{
		AppModuleBasic: AppModuleBasic{},
		accountKeeper:  accountKeeper,
	})
}

// module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	moduleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, moduleCdc, am.accountKeeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.accountKeeper)
	return moduleCdc.MustMarshalJSON(gs)
}
