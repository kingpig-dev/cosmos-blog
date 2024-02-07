package keeper

import (
	"context"

	"blog/x/blog/types"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TimedoutPostAll(ctx context.Context, req *types.QueryAllTimedoutPostRequest) (*types.QueryAllTimedoutPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timedoutPosts []types.TimedoutPost

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	timedoutPostStore := prefix.NewStore(store, types.KeyPrefix(types.TimedoutPostKey))

	pageRes, err := query.Paginate(timedoutPostStore, req.Pagination, func(key []byte, value []byte) error {
		var timedoutPost types.TimedoutPost
		if err := k.cdc.Unmarshal(value, &timedoutPost); err != nil {
			return err
		}

		timedoutPosts = append(timedoutPosts, timedoutPost)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimedoutPostResponse{TimedoutPost: timedoutPosts, Pagination: pageRes}, nil
}

func (k Keeper) TimedoutPost(ctx context.Context, req *types.QueryGetTimedoutPostRequest) (*types.QueryGetTimedoutPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	timedoutPost, found := k.GetTimedoutPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimedoutPostResponse{TimedoutPost: timedoutPost}, nil
}
