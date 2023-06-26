// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"errors"

	"github.com/lasthyphen/dijetalgo/codec"
	"github.com/lasthyphen/dijetalgo/database"
	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/snow"
	"github.com/lasthyphen/dijetalgo/vms/components/djtx"
	"github.com/lasthyphen/dijetalgo/vms/components/verify"
)

var (
	errNilTx = errors.New("nil tx is not valid")

	_ UnsignedTx = &BaseTx{}
)

// BaseTx is the basis of all transactions.
type BaseTx struct {
	djtx.BaseTx `serialize:"true"`
}

// Init sets the FxID fields in the inputs and outputs of this [BaseTx]
// Also sets the [ctx] in the OutputOwners to the given [vm.ctx] so that the
// addresses can be json marshalled into human readable format
func (t *BaseTx) Init(vm *VM) error {
	for _, in := range t.Ins {
		fx, err := vm.getParsedFx(in.In)
		if err != nil {
			return err
		}
		in.FxID = fx.ID
	}

	for _, out := range t.Outs {
		fx, err := vm.getParsedFx(out.Out)
		if err != nil {
			return err
		}
		out.FxID = fx.ID
		out.InitCtx(vm.ctx)
	}
	return nil
}

// SyntacticVerify that this transaction is well-formed.
func (t *BaseTx) SyntacticVerify(
	ctx *snow.Context,
	c codec.Manager,
	txFeeAssetID ids.ID,
	txFee uint64,
	_ uint64,
	_ int,
) error {
	if t == nil {
		return errNilTx
	}
	if err := t.MetadataVerify(ctx); err != nil {
		return err
	}

	return djtx.VerifyTx(
		txFee,
		txFeeAssetID,
		[][]*djtx.TransferableInput{t.Ins},
		[][]*djtx.TransferableOutput{t.Outs},
		c,
	)
}

// SemanticVerify that this transaction is valid to be spent.
func (t *BaseTx) SemanticVerify(vm *VM, tx UnsignedTx, creds []verify.Verifiable) error {
	for i, in := range t.Ins {
		cred := creds[i]
		if err := vm.verifyTransfer(tx, in, cred); err != nil {
			return err
		}
	}
	for _, out := range t.Outs {
		fxIndex, err := vm.getFx(out.Out)
		if err != nil {
			return err
		}
		if assetID := out.AssetID(); !vm.verifyFxUsage(fxIndex, assetID) {
			return errIncompatibleFx
		}
	}
	return nil
}

// ExecuteWithSideEffects writes the batch with any additional side effects
func (t *BaseTx) ExecuteWithSideEffects(_ *VM, batch database.Batch) error { return batch.Write() }
