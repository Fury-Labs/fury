package v5_0_0 //nolint:revive,stylecheck

import (
	lendkeeper "github.com/Fury-Labs/fury/x/lend/keeper"
	liquidationkeeper "github.com/Fury-Labs/fury/x/liquidation/keeper"
	vaultkeeper "github.com/Fury-Labs/fury/x/vault/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandlerV5Beta(
	mm *module.Manager,
	configurator module.Configurator,
	lk lendkeeper.Keeper,
	liqk liquidationkeeper.Keeper,
	vaultkeeper vaultkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		vm, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}
		SetVaultLengthCounter(ctx, vaultkeeper)
		err = FuncMigrateLiquidatedBorrow(ctx, lk, liqk)
		if err != nil {
			return nil, err
		}
		return vm, err
	}
}
