package keeper

import (
	"context"
	"encoding/binary"

	"blog/x/blog/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetSentPostCount get the total number of sentPost
func (k Keeper) GetSentPostCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SentPostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentPostCount set the total number of sentPost
func (k Keeper) SetSentPostCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SentPostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentPost appends a sentPost in the store with a new id and update the count
func (k Keeper) AppendSentPost(
	ctx context.Context,
	sentPost types.SentPost,
) uint64 {
	// Create the sentPost
	count := k.GetSentPostCount(ctx)

	// Set the ID of the appended value
	sentPost.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SentPostKey))
	appendedValue := k.cdc.MustMarshal(&sentPost)
	store.Set(GetSentPostIDBytes(sentPost.Id), appendedValue)

	// Update sentPost count
	k.SetSentPostCount(ctx, count+1)

	return count
}

// SetSentPost set a specific sentPost in the store
func (k Keeper) SetSentPost(ctx context.Context, sentPost types.SentPost) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SentPostKey))
	b := k.cdc.MustMarshal(&sentPost)
	store.Set(GetSentPostIDBytes(sentPost.Id), b)
}

// GetSentPost returns a sentPost from its id
func (k Keeper) GetSentPost(ctx context.Context, id uint64) (val types.SentPost, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SentPostKey))
	b := store.Get(GetSentPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentPost removes a sentPost from the store
func (k Keeper) RemoveSentPost(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SentPostKey))
	store.Delete(GetSentPostIDBytes(id))
}

// GetAllSentPost returns all sentPost
func (k Keeper) GetAllSentPost(ctx context.Context) (list []types.SentPost) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SentPostKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentPostIDBytes returns the byte representation of the ID
func GetSentPostIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.SentPostKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
