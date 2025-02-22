package keeper

import (
	"encoding/binary"
	"fmt"

	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/nft/internal/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the NFT Querier
const (
	QuerySupply       = "supply"
	QueryOwner        = "owner"
	QueryOwnerByDenom = "ownerByDenom"
	QueryCollection   = "collection"
	QueryDenoms       = "denoms"
	QueryNFT          = "nft"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QuerySupply:
			return querySupply(ctx, path[1:], req, k)
		case QueryOwner:
			return queryOwner(ctx, path[1:], req, k)
		case QueryOwnerByDenom:
			return queryOwnerByDenom(ctx, path[1:], req, k)
		case QueryCollection:
			return queryCollection(ctx, path[1:], req, k)
		case QueryNFT:
			return queryNFT(ctx, path[1:], req, k)
		case QueryDenoms:
			return queryDenoms(ctx, path[1:], req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nft query endpoint")
		}
	}
}

func queryDenoms(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	denoms := k.GetDenoms(ctx)
	bz := types.ModuleCdc.MustMarshalJSON(denoms)
	return bz, nil
}

func querySupply(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {

	var params types.QueryCollectionParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	collection, found := k.GetCollection(ctx, params.Denom)
	if !found {
		return nil, types.ErrUnknownCollection(types.DefaultCodespace, fmt.Sprintf("unknown denom %s", params.Denom))
	}

	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, uint64(collection.Supply()))
	return bz, nil
}

func queryOwnerByDenom(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryBalanceParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	var owner types.Owner

	var idCollections []types.IDCollection
	idCollection, _ := k.GetOwnerByDenom(ctx, params.Owner, params.Denom)
	owner.Address = params.Owner
	owner.IDCollections = append(idCollections, idCollection)

	bz := types.ModuleCdc.MustMarshalJSON(owner)
	return bz, nil
}

func queryOwner(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryBalanceParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}
	owner := k.GetOwner(ctx, params.Owner)
	bz := types.ModuleCdc.MustMarshalJSON(owner)
	return bz, nil
}

func queryCollection(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {

	var params types.QueryCollectionParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	collection, found := k.GetCollection(ctx, params.Denom)
	if !found {
		return nil, types.ErrUnknownCollection(types.DefaultCodespace, fmt.Sprintf("unknown denom %s", params.Denom))
	}

	// TODO: use the custom collection MarshalJSON

	// collections := types.NewCollections(collection)

	bz := types.ModuleCdc.MustMarshalJSON(collection)
	// bz, err := collections.MarshalJSON()
	// if err != nil {
	// 	return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	// }

	// var collectionBack types.Collection
	// types.ModuleCdc.MustUnmarshalJSON(bz, &collectionBack)
	// fmt.Println("collectionBack", collectionBack)

	return bz, nil
}

func queryNFT(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {

	var params types.QueryNFTParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	nft, err := k.GetNFT(ctx, params.Denom, params.TokenID)
	if err != nil {
		return nil, types.ErrUnknownNFT(types.DefaultCodespace, fmt.Sprintf("invalid NFT #%s from collection %s", params.TokenID, params.Denom))
	}

	bz := types.ModuleCdc.MustMarshalJSON(nft)
	return bz, nil
}
