// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/lasthyphen/dijetalgo/codec"
	"github.com/lasthyphen/dijetalgo/codec/linearcodec"
	"github.com/lasthyphen/dijetalgo/utils/units"
	"github.com/lasthyphen/dijetalgo/utils/wrappers"
)

const (
	// maxSize is the maximum allowed vertex size. It is necessary to deter DoS
	maxSize = units.MiB

	// noEpochTransitionsCodecVersion is the codec version that was used when
	// there were no epoch transitions
	noEpochTransitionsCodecVersion = uint16(0)

	// apricotCodecVersion is the codec version that was used when we added
	// epoch transitions
	apricotCodecVersion = uint16(1)
)

var c codec.Manager

func init() {
	codecV0 := linearcodec.New("serializeV0", maxSize)
	codecV1 := linearcodec.New("serializeV1", maxSize)
	c = codec.NewManager(maxSize)

	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterCodec(noEpochTransitionsCodecVersion, codecV0),
		c.RegisterCodec(apricotCodecVersion, codecV1),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
