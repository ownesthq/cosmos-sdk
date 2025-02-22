package store

import (
	"bytes"

	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/ownesthq/cosmos-sdk/store/types"
)

// Gets the first item.
func First(st KVStore, start, end []byte) (kv cmn.KVPair, ok bool) {
	iter := st.Iterator(start, end)
	if !iter.Valid() {
		return kv, false
	}
	defer iter.Close()

	return cmn.KVPair{Key: iter.Key(), Value: iter.Value()}, true
}

// Gets the last item.  `end` is exclusive.
func Last(st KVStore, start, end []byte) (kv cmn.KVPair, ok bool) {
	iter := st.ReverseIterator(end, start)
	if !iter.Valid() {
		if v := st.Get(start); v != nil {
			return cmn.KVPair{Key: types.Cp(start), Value: types.Cp(v)}, true
		}
		return kv, false
	}
	defer iter.Close()

	if bytes.Equal(iter.Key(), end) {
		// Skip this one, end is exclusive.
		iter.Next()
		if !iter.Valid() {
			return kv, false
		}
	}

	return cmn.KVPair{Key: iter.Key(), Value: iter.Value()}, true
}
