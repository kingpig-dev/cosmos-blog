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

func (k Keeper) SentPostAll(ctx context.Context, req *types.QueryAllSentPostRequest) (*types.QueryAllSentPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentPosts []types.SentPost

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	sentPostStore := prefix.NewStore(store, types.KeyPrefix(types.SentPostKey))

	pageRes, err := query.Paginate(sentPostStore, req.Pagination, func(key []byte, value []byte) error {
		var sentPost types.SentPost
		if err := k.cdc.Unmarshal(value, &sentPost); err != nil {
			return err
		}

		sentPosts = append(sentPosts, sentPost)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentPostResponse{SentPost: sentPosts, Pagination: pageRes}, nil
}

func (k Keeper) SentPost(ctx context.Context, req *types.QueryGetSentPostRequest) (*types.QueryGetSentPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sentPost, found := k.GetSentPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentPostResponse{SentPost: sentPost}, nil
}
