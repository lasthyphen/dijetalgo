// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/snow"
	"github.com/lasthyphen/dijetalgo/vms/components/djtx"
	"github.com/lasthyphen/dijetalgo/vms/components/verify"
	"github.com/lasthyphen/dijetalgo/vms/nftfx"
	"github.com/lasthyphen/dijetalgo/vms/propertyfx"
	"github.com/lasthyphen/dijetalgo/vms/secp256k1fx"
)

var (
	_ Fx = &secp256k1fx.Fx{}
	_ Fx = &nftfx.Fx{}
	_ Fx = &propertyfx.Fx{}
)

type parsedFx struct {
	ID ids.ID
	Fx Fx
}

// Fx is the interface a feature extension must implement to support the AVM.
type Fx interface {
	// Initialize this feature extension to be running under this VM. Should
	// return an error if the VM is incompatible.
	Initialize(vm interface{}) error

	// Notify this Fx that the VM is in bootstrapping
	Bootstrapping() error

	// Notify this Fx that the VM is bootstrapped
	Bootstrapped() error

	// VerifyTransfer verifies that the specified transaction can spend the
	// provided utxo with no restrictions on the destination. If the transaction
	// can't spend the output based on the input and credential, a non-nil error
	// should be returned.
	VerifyTransfer(tx, in, cred, utxo interface{}) error

	// VerifyOperation verifies that the specified transaction can spend the
	// provided utxos conditioned on the result being restricted to the provided
	// outputs. If the transaction can't spend the output based on the input and
	// credential, a non-nil error  should be returned.
	VerifyOperation(tx, op, cred interface{}, utxos []interface{}) error
}

type FxOperation interface {
	verify.Verifiable
	snow.ContextInitializable
	djtx.Coster

	Outs() []verify.State
}

type FxCredential struct {
	FxID              ids.ID `serialize:"false" json:"fxID"`
	verify.Verifiable `serialize:"true" json:"credential"`
}
