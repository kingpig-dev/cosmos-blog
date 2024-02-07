package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "blog/testutil/keeper"
	"blog/testutil/nullify"
	"blog/x/blog/types"
)

func TestSentPostQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	msgs := createNSentPost(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSentPostRequest
		response *types.QueryGetSentPostResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSentPostRequest{Id: msgs[0].Id},
			response: &types.QueryGetSentPostResponse{SentPost: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetSentPostRequest{Id: msgs[1].Id},
			response: &types.QueryGetSentPostResponse{SentPost: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetSentPostRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SentPost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestSentPostQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	msgs := createNSentPost(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSentPostRequest {
		return &types.QueryAllSentPostRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SentPostAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentPost), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SentPost),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SentPostAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentPost), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SentPost),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SentPostAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SentPost),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SentPostAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
