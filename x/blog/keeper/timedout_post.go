package keeper

import (
	"context"
	"encoding/binary"

	"blog/x/blog/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetTimedoutPostCount get the total number of timedoutPost
func (k Keeper) GetTimedoutPostCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TimedoutPostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutPostCount set the total number of timedoutPost
func (k Keeper) SetTimedoutPostCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TimedoutPostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutPost appends a timedoutPost in the store with a new id and update the count
func (k Keeper) AppendTimedoutPost(
	ctx context.Context,
	timedoutPost types.TimedoutPost,
) uint64 {
	// Create the timedoutPost
	count := k.GetTimedoutPostCount(ctx)

	// Set the ID of the appended value
	timedoutPost.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TimedoutPostKey))
	appendedValue := k.cdc.MustMarshal(&timedoutPost)
	store.Set(GetTimedoutPostIDBytes(timedoutPost.Id), appendedValue)

	// Update timedoutPost count
	k.SetTimedoutPostCount(ctx, count+1)

	return count
}

// SetTimedoutPost set a specific timedoutPost in the store
func (k Keeper) SetTimedoutPost(ctx context.Context, timedoutPost types.TimedoutPost) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TimedoutPostKey))
	b := k.cdc.MustMarshal(&timedoutPost)
	store.Set(GetTimedoutPostIDBytes(timedoutPost.Id), b)
}

// GetTimedoutPost returns a timedoutPost from its id
func (k Keeper) GetTimedoutPost(ctx context.Context, id uint64) (val types.TimedoutPost, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TimedoutPostKey))
	b := store.Get(GetTimedoutPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutPost removes a timedoutPost from the store
func (k Keeper) RemoveTimedoutPost(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TimedoutPostKey))
	store.Delete(GetTimedoutPostIDBytes(id))
}

// GetAllTimedoutPost returns all timedoutPost
func (k Keeper) GetAllTimedoutPost(ctx context.Context) (list []types.TimedoutPost) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TimedoutPostKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutPostIDBytes returns the byte representation of the ID
func GetTimedoutPostIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.TimedoutPostKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
