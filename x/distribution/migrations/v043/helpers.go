package v043

import (
	"github.com/opzlabs/cosmos-sdk/store/prefix"
	sdk "github.com/opzlabs/cosmos-sdk/types"
	"github.com/opzlabs/cosmos-sdk/types/address"
	v042auth "github.com/opzlabs/cosmos-sdk/x/auth/migrations/v042"
)

// MigratePrefixAddress is a helper function that migrates all keys of format:
// prefix_bytes | address_bytes
// into format:
// prefix_bytes | address_len (1 byte) | address_bytes
func MigratePrefixAddress(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		addr := oldStoreIter.Key()
		newStoreKey := prefixBz
		newStoreKey = append(newStoreKey, address.MustLengthPrefix(addr)...)

		// Set new key on store. Values don't change.
		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}

// MigratePrefixAddressBytes is a helper function that migrates all keys of format:
// prefix_bytes | address_bytes | arbitrary_bytes
// into format:
// prefix_bytes | address_len (1 byte) | address_bytes | arbitrary_bytes
func MigratePrefixAddressBytes(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		addr := oldStoreIter.Key()[:v042auth.AddrLen]
		endBz := oldStoreIter.Key()[v042auth.AddrLen:]
		newStoreKey := append(append(prefixBz, address.MustLengthPrefix(addr)...), endBz...)

		// Set new key on store. Values don't change.
		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}

// MigratePrefixAddressAddress is a helper function that migrates all keys of format:
// prefix_bytes | address_1_bytes | address_2_bytes
// into format:
// prefix_bytes | address_1_len (1 byte) | address_1_bytes | address_2_len (1 byte) | address_2_bytes
func MigratePrefixAddressAddress(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		addr1 := oldStoreIter.Key()[:v042auth.AddrLen]
		addr2 := oldStoreIter.Key()[v042auth.AddrLen:]
		newStoreKey := append(append(prefixBz, address.MustLengthPrefix(addr1)...), address.MustLengthPrefix(addr2)...)

		// Set new key on store. Values don't change.
		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}
