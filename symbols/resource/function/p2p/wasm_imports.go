//go:build wasi || wasm
// +build wasi wasm

package functionSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/smartops
//export getFunctionP2PName
func GetName(resourceId uint32, data *byte) (error errno.Error)

//go:wasm-module taubyte/smartops
//export getFunctionP2PNameSize
func GetNameSize(resourceId uint32, size *uint32) (error errno.Error)
