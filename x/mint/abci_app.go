package mint

import (
	"github.com/cosmos/cosmos-sdk"
)

// mint new tokens for the previous block
func BeginBlocker(ctx sdk.Context, k Keeper) {

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	totalSupply := k.sk.TotalTokens(ctx)
	bondedRatio := k.sk.BondedRatio(ctx)
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalSupply)
	k.SetMinter(ctx, minter)

	// mint coins, add to collected fees, update supply
	mintedCoin := minter.BlockProvision(params)
	k.fck.AddCollectedFees(ctx, sdk.Coins{mintedCoin})
	k.sk.InflateSupply(ctx, mintedCoin.Amount)

}
