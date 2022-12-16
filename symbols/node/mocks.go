//go:build !wasi && !wasm
// +build !wasi,!wasm

package nodeSym

import (
	"unsafe"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/ipfs/client"
)

type MockData struct {
	Cid cid.Cid
}

func (_m MockData) Mock() *MockData {
	m := &_m

	m.Id()

	return m
}

func (m *MockData) Id() {
	GetNodeId = func(cidPtr *byte) (error errno.Error) {
		cidData := m.Cid.Bytes()
		_cidPtr := unsafe.Slice(cidPtr, client.CidBufferSize)
		copy(_cidPtr, cidData)

		return 0
	}
}
