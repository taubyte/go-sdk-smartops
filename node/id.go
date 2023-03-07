package node

import (
	"fmt"

	"github.com/ipfs/go-cid"
	nodeSym "github.com/taubyte/go-sdk-smartops/symbols/node"
	"github.com/taubyte/go-sdk/utils/codec"
)

func Id() (cid.Cid, error) {
	_cid := codec.CidReader()
	err0 := nodeSym.GetNodeId(_cid.Ptr())
	if err0 != 0 {
		return cid.Cid{}, fmt.Errorf("getting cid failed with: %s", err0)
	}

	return _cid.Parse()
}
