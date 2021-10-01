package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sifapp "github.com/Sifchain/sifnode/app"
)

func TestKeeper_GetCrossChainFee(t *testing.T) {
	app := sifapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// test against wrong network identity
	networkDescriptor := types.NewNetworkIdentity(types.NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED)
	token := "ceth"
	_, err := app.OracleKeeper.GetCrossChainFee(ctx, networkDescriptor)
	assert.Error(t, err)

	// test if configure not set for ethereum
	networkDescriptor = types.NewNetworkIdentity(types.NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM)
	_, err = app.OracleKeeper.GetCrossChainFee(ctx, networkDescriptor)
	assert.Error(t, err)

	gas := sdk.NewInt(0)
	lockCost := sdk.NewInt(0)
	burnCost := sdk.NewInt(0)
	firstLockDoublePeggyCost := sdk.NewInt(0)

	app.OracleKeeper.SetCrossChainFee(ctx, networkDescriptor, token, gas, lockCost, burnCost, firstLockDoublePeggyCost)

	// case for well set the configure
	tokenStored, err := app.OracleKeeper.GetCrossChainFee(ctx, networkDescriptor)
	assert.NoError(t, err)
	assert.Equal(t, token, tokenStored)
}

func TestKeeper_SetCrossChainFee(t *testing.T) {
	app := sifapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	networkDescriptor := types.NewNetworkIdentity(types.NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM)
	token := "ceth"
	gas := sdk.NewInt(0)
	lockCost := sdk.NewInt(0)
	burnCost := sdk.NewInt(0)
	firstLockDoublePeggyCost := sdk.NewInt(0)

	app.OracleKeeper.SetCrossChainFee(ctx, networkDescriptor, token, gas, lockCost, burnCost, firstLockDoublePeggyCost)

	tokenStored, err := app.OracleKeeper.GetCrossChainFee(ctx, networkDescriptor)
	assert.NoError(t, err)
	assert.Equal(t, token, tokenStored)
}
