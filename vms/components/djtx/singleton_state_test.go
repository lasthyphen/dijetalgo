// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package djtx

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lasthyphen/dijetalgo/database/memdb"
)

func TestSingletonState(t *testing.T) {
	assert := assert.New(t)

	db := memdb.New()
	s := NewSingletonState(db)

	isInitialized, err := s.IsInitialized()
	assert.NoError(err)
	assert.False(isInitialized)

	err = s.SetInitialized()
	assert.NoError(err)

	isInitialized, err = s.IsInitialized()
	assert.NoError(err)
	assert.True(isInitialized)
}
