// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"time"

	"github.com/lasthyphen/dijetalgo/ids"
)

type PeerID struct {
	IP           string    `json:"ip"`
	PublicIP     string    `json:"publicIP,omitempty"`
	ID           string    `json:"nodeID"`
	Version      string    `json:"version"`
	LastSent     time.Time `json:"lastSent"`
	LastReceived time.Time `json:"lastReceived"`
	Benched      []ids.ID  `json:"benched"`
}
