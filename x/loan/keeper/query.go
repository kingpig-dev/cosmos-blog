package keeper

import (
	"blog/x/loan/types"
)

var _ types.QueryServer = Keeper{}
