//go:build !wasi && !wasm
// +build !wasi,!wasm

package nodeSym

import "github.com/taubyte/go-sdk/errno"

var GetNodeId = func(cidPtr *byte) (error errno.Error) {
	return 0
}
