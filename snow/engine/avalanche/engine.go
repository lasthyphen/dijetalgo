// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/snow/consensus/avalanche"
	"github.com/lasthyphen/dijetalgo/snow/engine/common"
)

// Engine describes the events that can occur on a consensus instance
type Engine interface {
	common.Engine

	// Initialize this engine.
	Initialize(Config) error

	// GetVtx returns a vertex by its ID.
	// Returns an error if unknown.
	GetVtx(vtxID ids.ID) (avalanche.Vertex, error)
}
