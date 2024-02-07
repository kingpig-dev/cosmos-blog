package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/keeper"
	"blog/x/blog/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)
	// Set all the loan
	for _, elem := range genState.LoanList {
		k.SetLoan(ctx, elem)
	}

	// Set loan count
	k.SetLoanCount(ctx, genState.LoanCount)
	// Set all the sentPost
	for _, elem := range genState.SentPostList {
		k.SetSentPost(ctx, elem)
	}

	// Set sentPost count
	k.SetSentPostCount(ctx, genState.SentPostCount)
	// Set all the timedoutPost
	for _, elem := range genState.TimedoutPostList {
		k.SetTimedoutPost(ctx, elem)
	}

	// Set timedoutPost count
	k.SetTimedoutPostCount(ctx, genState.TimedoutPostCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PostList = k.GetAllPost(ctx)
	genesis.PostCount = k.GetPostCount(ctx)
	genesis.LoanList = k.GetAllLoan(ctx)
	genesis.LoanCount = k.GetLoanCount(ctx)
	genesis.SentPostList = k.GetAllSentPost(ctx)
	genesis.SentPostCount = k.GetSentPostCount(ctx)
	genesis.TimedoutPostList = k.GetAllTimedoutPost(ctx)
	genesis.TimedoutPostCount = k.GetTimedoutPostCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
