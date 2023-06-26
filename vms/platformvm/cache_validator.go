// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"github.com/lasthyphen/dijetalgo/ids"
)

var _ validator = &validatorImpl{}

type validator interface {
	Delegators() []*UnsignedAddDelegatorTx
	SubnetValidators() map[ids.ID]*UnsignedAddSubnetValidatorTx
}

type validatorImpl struct {
	// sorted in order of next operation, either addition or removal.
	delegators []*UnsignedAddDelegatorTx
	// maps subnetID to tx
	subnets map[ids.ID]*UnsignedAddSubnetValidatorTx
}

func (v *validatorImpl) Delegators() []*UnsignedAddDelegatorTx {
	return v.delegators
}

func (v *validatorImpl) SubnetValidators() map[ids.ID]*UnsignedAddSubnetValidatorTx {
	return v.subnets
}
