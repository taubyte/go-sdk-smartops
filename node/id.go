package node

import (
	"fmt"

	"github.com/ipfs/go-cid"
	nodeSym "github.com/taubyte/go-sdk-smartops/symbols/node"
	"github.com/taubyte/go-sdk/ipfs/client"
)

func Id() (cid.Cid, error) {
	_cid := make([]byte, client.CidBufferSize)

	err0 := nodeSym.GetNodeId(&_cid[0])
	if err0 != 0 {
		return cid.Cid{}, fmt.Errorf("Failed getting cid with %v", err0)
	}

	_, cidFromBytes, err := cid.CidFromBytes(_cid)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("Failed decoding cid with: %v", err)
	}

	return cidFromBytes, err
}
