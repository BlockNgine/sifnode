package keeper

import (
	"errors"
	"fmt"

	oracletypes "github.com/Sifchain/sifnode/x/oracle/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Sifchain/sifnode/x/ethbridge/types"
)

const errorMessageKey = "errorMessageKey"

// Keeper maintains the link to data storage and
// exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc codec.BinaryMarshaler // The wire codec for binary encoding/decoding.

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	oracleKeeper  types.OracleKeeper
	storeKey      sdk.StoreKey
}

// GetBankKeeper
func (k Keeper) GetBankKeeper() types.BankKeeper {
	return k.bankKeeper
}

// NewKeeper creates new instances of the oracle Keeper
func NewKeeper(cdc codec.BinaryMarshaler, bankKeeper types.BankKeeper, oracleKeeper types.OracleKeeper, accountKeeper types.AccountKeeper, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		oracleKeeper:  oracleKeeper,
		storeKey:      storeKey,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ProcessClaim processes a new claim coming in from a validator
func (k Keeper) ProcessClaim(ctx sdk.Context, claim *types.EthBridgeClaim) (oracletypes.StatusText, error) {
	return k.oracleKeeper.ProcessClaim(ctx, claim.NetworkDescriptor, claim.GetProphecyID(), claim.ValidatorAddress)
}

// ProcessSuccessfulClaim processes a claim that has just completed successfully with consensus
func (k Keeper) ProcessSuccessfulClaim(ctx sdk.Context, claim *types.EthBridgeClaim) error {
	logger := k.Logger(ctx)

	var coins sdk.Coins
	var err error
	switch claim.ClaimType {
	case types.ClaimType_CLAIM_TYPE_LOCK:
		symbol := fmt.Sprintf("%v%v", types.PeggedCoinPrefix, claim.Symbol)
		k.AddPeggyToken(ctx, symbol)

		coins = sdk.Coins{sdk.NewCoin(symbol, claim.Amount)}
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	case types.ClaimType_CLAIM_TYPE_BURN:
		coins = sdk.Coins{sdk.NewCoin(claim.Symbol, claim.Amount)}
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	default:
		err = types.ErrInvalidClaimType
	}

	if err != nil {
		logger.Error("failed to process successful claim.",
			errorMessageKey, err.Error())
		return err
	}

	receiverAddress, err := sdk.AccAddressFromBech32(claim.CosmosReceiver)

	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiverAddress, coins,
	); err != nil {
		panic(err)
	}

	return nil
}

// ProcessBurn processes the burn of bridged coins from the given sender
func (k Keeper) ProcessBurn(ctx sdk.Context, cosmosSender sdk.AccAddress, msg *types.MsgBurn) error {
	logger := k.Logger(ctx)
	var coins sdk.Coins
	networkIdentity := oracletypes.NewNetworkIdentity(msg.NetworkDescriptor)
	nativeTokenConfig, err := k.oracleKeeper.GetNativeTokenConfig(ctx, networkIdentity)

	if err != nil {
		return err
	}

	minimumBurn := nativeTokenConfig.MinimumBurnCost.Mul(nativeTokenConfig.NativeGas)
	if msg.NativeTokenAmount.LT(minimumBurn) {
		return errors.New("native token amount in message less than minimum burn")
	}

	if k.IsNativeTokenReceiverAccountSet(ctx) {
		coins = sdk.NewCoins(sdk.NewCoin(nativeTokenConfig.NativeToken, msg.NativeTokenAmount))

		err := k.bankKeeper.SendCoins(ctx, cosmosSender, k.GetNativeTokenReceiverAccount(ctx), coins)
		if err != nil {
			logger.Error("failed to send native_token from account to account.",
				errorMessageKey, err.Error())
			return err
		}

		coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount))

	} else {
		if msg.Symbol == nativeTokenConfig.NativeToken {
			coins = sdk.NewCoins(sdk.NewCoin(nativeTokenConfig.NativeToken, msg.NativeTokenAmount.Add(msg.Amount)))
		} else {
			coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount), sdk.NewCoin(nativeTokenConfig.NativeToken, msg.NativeTokenAmount))
		}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, cosmosSender, types.ModuleName, coins)
	if err != nil {
		logger.Error("failed to send native_token from module to account.",
			errorMessageKey, err.Error())
		return err
	}

	coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount))
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		logger.Error("failed to burn locked coin.",
			errorMessageKey, err.Error())
		return err
	}

	return nil
}

// ProcessLock processes the lockup of cosmos coins from the given sender
func (k Keeper) ProcessLock(ctx sdk.Context, cosmosSender sdk.AccAddress, msg *types.MsgLock) error {
	logger := k.Logger(ctx)
	var coins sdk.Coins
	networkIdentity := oracletypes.NewNetworkIdentity(msg.NetworkDescriptor)
	nativeTokenConfig, err := k.oracleKeeper.GetNativeTokenConfig(ctx, networkIdentity)

	if err != nil {
		return err
	}

	minimumLock := nativeTokenConfig.MinimumLockCost.Mul(nativeTokenConfig.NativeGas)
	if msg.NativeTokenAmount.LT(minimumLock) {
		return errors.New("native token amount in message less than minimum lock")
	}

	if k.IsNativeTokenReceiverAccountSet(ctx) {
		coins = sdk.NewCoins(sdk.NewCoin(nativeTokenConfig.NativeToken, msg.NativeTokenAmount))

		err := k.bankKeeper.SendCoins(ctx, cosmosSender, k.GetNativeTokenReceiverAccount(ctx), coins)
		if err != nil {
			logger.Error("failed to send native_token from account to account.",
				errorMessageKey, err.Error())
			return err
		}

		coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount))

	} else {
		coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount), sdk.NewCoin(nativeTokenConfig.NativeToken, msg.NativeTokenAmount))
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, cosmosSender, types.ModuleName, coins)

	if err != nil {
		logger.Error("failed to transfer coin from account to module.",
			errorMessageKey, err.Error())
		return err
	}

	coins = sdk.NewCoins(sdk.NewCoin(msg.Symbol, msg.Amount))
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		logger.Error("failed to burn burned coin.",
			errorMessageKey, err.Error())
		return err
	}
	return nil
}

// ProcessUpdateWhiteListValidator processes the update whitelist validator from admin
func (k Keeper) ProcessUpdateWhiteListValidator(ctx sdk.Context, networkDescriptor oracletypes.NetworkDescriptor, cosmosSender sdk.AccAddress, validator sdk.ValAddress, power uint32) error {
	return k.oracleKeeper.ProcessUpdateWhiteListValidator(ctx, networkDescriptor, cosmosSender, validator, power)
}

// ProcessUpdateNativeTokenReceiverAccount processes the update whitelist validator from admin
func (k Keeper) ProcessUpdateNativeTokenReceiverAccount(ctx sdk.Context, cosmosSender sdk.AccAddress, nativeTokenReceiverAccount sdk.AccAddress) error {
	logger := k.Logger(ctx)
	if !k.oracleKeeper.IsAdminAccount(ctx, cosmosSender) {
		logger.Error("cosmos sender is not admin account.")
		return errors.New("only admin account can update NativeToken receiver account")
	}

	k.SetNativeTokenReceiverAccount(ctx, nativeTokenReceiverAccount)
	return nil
}

// ProcessRescueNativeToken transfer NativeToken from ethbridge module to an account
func (k Keeper) ProcessRescueNativeToken(ctx sdk.Context, msg *types.MsgRescueNativeToken) error {
	logger := k.Logger(ctx)

	cosmosSender, err := sdk.AccAddressFromBech32(msg.CosmosSender)
	if err != nil {
		return err
	}

	cosmosReceiver, err := sdk.AccAddressFromBech32(msg.CosmosReceiver)
	if err != nil {
		return err
	}

	if !k.oracleKeeper.IsAdminAccount(ctx, cosmosSender) {
		logger.Error("cosmos sender is not admin account.")
		return errors.New("only admin account can call rescue NativeToken")
	}

	coins := sdk.NewCoins(sdk.NewCoin(msg.NativeTokenSymbol, msg.NativeTokenAmount))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, cosmosReceiver, coins)

	if err != nil {
		logger.Error("failed to transfer coin from module to account.",
			errorMessageKey, err.Error())
		return err
	}
	return nil
}

// ProcessSetNativeToken processes the set native token from admin
func (k Keeper) ProcessSetNativeToken(ctx sdk.Context, msg *types.MsgSetNativeToken) error {
	logger := k.Logger(ctx)

	cosmosSender, err := sdk.AccAddressFromBech32(msg.CosmosSender)
	if err != nil {
		return err
	}

	if !k.oracleKeeper.IsAdminAccount(ctx, cosmosSender) {
		logger.Error("cosmos sender is not admin account.")
		return errors.New("only admin account can set native token")
	}
	return k.oracleKeeper.ProcessSetNativeToken(ctx, msg.NetworkDescriptor, msg.NativeToken, msg.NativeGas, msg.MinimumBurnCost, msg.MinimumLockCost)
}

// Exists chec if the key existed in db.
func (k Keeper) Exists(ctx sdk.Context, key []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(key)
}
