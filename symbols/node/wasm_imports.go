//go:build wasi || wasm
// +build wasi wasm

package nodeSym

import "github.com/taubyte/go-sdk/errno"

//go:wasm-module taubyte/smartops
//export getNodeId
func GetNodeId(cidPtr *byte) (error errno.Error)
