package evm

import (
	anteinterfaces "github.com/cosmos/evm/ante/interfaces"
	evmante "github.com/cosmos/evm/x/vm/ante"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

var _ sdktypes.AnteDecorator = &EthSetupContextDecorator{}

// EthSetupContextDecorator is adapted from SetUpContextDecorator from cosmos-sdk, it ignores gas consumption
// by setting the gas meter to infinite
type EthSetupContextDecorator struct {
	evmKeeper anteinterfaces.EVMKeeper
}

func NewEthSetUpContextDecorator(evmKeeper anteinterfaces.EVMKeeper) EthSetupContextDecorator {
	return EthSetupContextDecorator{
		evmKeeper: evmKeeper,
	}
}

func (esc EthSetupContextDecorator) AnteHandle(ctx sdktypes.Context, tx sdktypes.Tx, simulate bool, next sdktypes.AnteHandler) (newCtx sdktypes.Context, err error) {
	newCtx, err = SetupContextAndResetTransientGas(ctx, tx, esc.evmKeeper)
	if err != nil {
		return ctx, err
	}
	return next(newCtx, tx, simulate)
}

// SetupContextAndResetTransientGas modifies the context to be used in the
// execution of the ante handler associated with an EVM transaction. Previous
// gas consumed is reset in the transient store.
func SetupContextAndResetTransientGas(ctx sdktypes.Context, tx sdktypes.Tx, evmKeeper anteinterfaces.EVMKeeper) (sdktypes.Context, error) {
	// all transactions must implement GasTx
	_, ok := tx.(authante.GasTx)
	if !ok {
		return ctx, errorsmod.Wrapf(errortypes.ErrInvalidType, "invalid transaction type %T, expected GasTx", tx)
	}

	// To have gas consumption consistent with Ethereum, we need to:
	//     1. Set an empty gas config for both KV and transient store.
	//     2. Set an infinite gas meter.
	newCtx := evmante.BuildEvmExecutionCtx(ctx).
		WithGasMeter(storetypes.NewInfiniteGasMeter())

	// Reset transient gas used to prepare the execution of current cosmos tx.
	// Transient gas-used is necessary to sum the gas-used of cosmos tx, when it contains multiple eth msgs.
	//
	// TODO: add more context here to explain why gas used is reset. Not clear
	// from docstring.
	evmKeeper.ResetTransientGasUsed(ctx)

	return newCtx, nil
}
