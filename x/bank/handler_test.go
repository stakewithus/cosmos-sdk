package bank

import (
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res := h(sdk.Context{}, sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized bank message type"))
}
