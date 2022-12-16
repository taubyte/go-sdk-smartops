package node_test

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk-smartops/node"
	nodeSym "github.com/taubyte/go-sdk-smartops/symbols/node"
	"github.com/taubyte/go-sdk/errno"
)

func TestId(t *testing.T) {
	testStringCid := "bafkreibrl5n5w5wqpdcdxcwaazheualemevr7ttxzbutiw74stdvrfhn2m"

	_cid, _ := cid.Parse(testStringCid)
	m := nodeSym.MockData{
		Cid: _cid,
	}.Mock()

	testId, err := node.Id()
	if err != nil {
		t.Error(err)
		return
	}

	if testId.String() != testStringCid {
		t.Errorf("Expected %v, got %v", testStringCid, testId.String())
	}

	m.Cid = cid.Cid{}
	m.Mock()

	_, err = node.Id()
	if err == nil {
		t.Error("Expected error")
		return
	}

	nodeSym.GetNodeId = func(cidPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = node.Id()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
