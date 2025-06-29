package itrie

import (
	"testing"

	"github.com/lerax-chain/lerax-chain/state"
)

func TestState(t *testing.T) {
	state.TestState(t, buildPreState)
}

func buildPreState(pre state.PreStates) state.Snapshot {
	storage := NewMemoryStorage()
	st := NewState(storage)

	return st.NewSnapshot()
}
