// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsubnetlookup

import (
	"context"

	"github.com/lasthyphen/dijetalgo/ids"
	"github.com/lasthyphen/dijetalgo/snow"
	"github.com/lasthyphen/dijetalgo/vms/rpcchainvm/gsubnetlookup/gsubnetlookupproto"
)

var _ gsubnetlookupproto.SubnetLookupServer = &Server{}

// Server is a subnet lookup that is managed over RPC.
type Server struct {
	gsubnetlookupproto.UnimplementedSubnetLookupServer
	aliaser snow.SubnetLookup
}

// NewServer returns a subnet lookup connected to a remote subnet lookup
func NewServer(aliaser snow.SubnetLookup) *Server {
	return &Server{aliaser: aliaser}
}

func (s *Server) SubnetID(
	_ context.Context,
	req *gsubnetlookupproto.SubnetIDRequest,
) (*gsubnetlookupproto.SubnetIDResponse, error) {
	chainID, err := ids.ToID(req.ChainID)
	if err != nil {
		return nil, err
	}
	id, err := s.aliaser.SubnetID(chainID)
	if err != nil {
		return nil, err
	}
	return &gsubnetlookupproto.SubnetIDResponse{
		Id: id[:],
	}, nil
}
