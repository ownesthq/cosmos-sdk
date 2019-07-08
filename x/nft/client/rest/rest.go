package rest

import (
	"github.com/gorilla/mux"

	"github.com/ownesthq/cosmos-sdk/client/context"
	"github.com/ownesthq/cosmos-sdk/codec"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, queryRoute string) {
	registerQueryRoutes(cliCtx, r, cdc, queryRoute)
	registerTxRoutes(cliCtx, r, cdc, queryRoute)
}
