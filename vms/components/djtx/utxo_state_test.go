// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package djtx

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lasthyphen/dijetalgo/codec"
	"github.com/lasthyphen/dijetalgo/codec/linearcodec"
	"github.com/lasthyphen/dijetalgo/database"
	"github.com/lasthyphen/dijetalgo/database/memdb"
	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/utils/wrappers"
	"github.com/lasthyphen/dijetalgo/vms/secp256k1fx"
)

func TestUTXOState(t *testing.T) {
	assert := assert.New(t)
	txID := ids.GenerateTestID()
	assetID := ids.GenerateTestID()
	addr := ids.GenerateTestShortID()
	utxoID := ids.GenerateTestID()
	utxo := &UTXO{
		UTXOID: UTXOID{
			TxID:        txID,
			OutputIndex: 0,
		},
		Asset: Asset{ID: assetID},
		Out: &secp256k1fx.TransferOutput{
			Amt: 12345,
			OutputOwners: secp256k1fx.OutputOwners{
				Locktime:  54321,
				Threshold: 1,
				Addrs:     []ids.ShortID{addr},
			},
		},
	}

	c := linearcodec.NewDefault()
	manager := codec.NewDefaultManager()

	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&secp256k1fx.MintOutput{}),
		c.RegisterType(&secp256k1fx.TransferOutput{}),
		c.RegisterType(&secp256k1fx.Input{}),
		c.RegisterType(&secp256k1fx.TransferInput{}),
		c.RegisterType(&secp256k1fx.Credential{}),
		manager.RegisterCodec(codecVersion, c),
	)
	assert.NoError(errs.Err)

	db := memdb.New()
	s := NewUTXOState(db, manager)

	_, err := s.GetUTXO(utxoID)
	assert.Equal(database.ErrNotFound, err)

	_, err = s.GetUTXO(utxoID)
	assert.Equal(database.ErrNotFound, err)

	err = s.DeleteUTXO(utxoID)
	assert.Equal(database.ErrNotFound, err)

	err = s.PutUTXO(utxoID, utxo)
	assert.NoError(err)

	utxoIDs, err := s.UTXOIDs(addr[:], ids.Empty, 5)
	assert.NoError(err)
	assert.Equal([]ids.ID{utxoID}, utxoIDs)

	readUTXO, err := s.GetUTXO(utxoID)
	assert.NoError(err)
	assert.Equal(utxo, readUTXO)

	err = s.DeleteUTXO(utxoID)
	assert.NoError(err)

	_, err = s.GetUTXO(utxoID)
	assert.Equal(database.ErrNotFound, err)

	err = s.PutUTXO(utxoID, utxo)
	assert.NoError(err)

	s = NewUTXOState(db, manager)

	readUTXO, err = s.GetUTXO(utxoID)
	assert.NoError(err)
	assert.Equal(utxo, readUTXO)

	utxoIDs, err = s.UTXOIDs(addr[:], ids.Empty, 5)
	assert.NoError(err)
	assert.Equal([]ids.ID{utxoID}, utxoIDs)
}
