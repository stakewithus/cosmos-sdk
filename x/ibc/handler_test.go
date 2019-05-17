package ibc

import (
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	m := Mapper{}
	h := NewHandler(m, nil)

	res := h(sdk.Context{}, sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized IBC message type"))
}
