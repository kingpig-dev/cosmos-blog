package keeper

import (
	"context"
	"encoding/binary"

	"blog/x/blog/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetPostCount get the total number of post
func (k Keeper) GetPostCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPostCount set the total number of post
func (k Keeper) SetPostCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPost appends a post in the store with a new id and update the count
func (k Keeper) AppendPost(
	ctx context.Context,
	post types.Post,
) uint64 {
	// Create the post
	count := k.GetPostCount(ctx)

	// Set the ID of the appended value
	post.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), appendedValue)

	// Update post count
	k.SetPostCount(ctx, count+1)

	return count
}

// SetPost set a specific post in the store
func (k Keeper) SetPost(ctx context.Context, post types.Post) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), b)
}

// GetPost returns a post from its id
func (k Keeper) GetPost(ctx context.Context, id uint64) (val types.Post, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	b := store.Get(GetPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePost removes a post from the store
func (k Keeper) RemovePost(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	store.Delete(GetPostIDBytes(id))
}

// GetAllPost returns all post
func (k Keeper) GetAllPost(ctx context.Context) (list []types.Post) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPostIDBytes returns the byte representation of the ID
func GetPostIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PostKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
