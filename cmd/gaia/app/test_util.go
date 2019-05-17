package app

import (
	"io"

	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/x/staking"
	dbm "github.com/tendermint/tendermint/libs/db"
)

// used for debugging by gaia/cmd/gaiadebug
// NOTE to not use this function with non-test code
func NewGaiaAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) (gapp *GaiaApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	gapp = NewGaiaApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return gapp, gapp.keyMain, gapp.keyStaking, gapp.stakingKeeper
}
