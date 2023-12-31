// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/utils/constants"
	"github.com/lasthyphen/dijetalgo/utils/logging"
)

var _ Handler = NoopHandler{}

type Handler interface {
	HandleTx(nodeID ids.ShortID, requestID uint32, msg *Tx) error
}

type NoopHandler struct {
	Log logging.Logger
}

func (h NoopHandler) HandleTx(nodeID ids.ShortID, requestID uint32, _ *Tx) error {
	h.Log.Debug(
		"dropping unexpected Tx message from %s with requestID %s",
		nodeID.PrefixedString(constants.NodeIDPrefix),
		requestID,
	)
	return nil
}
