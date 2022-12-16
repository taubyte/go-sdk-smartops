//go:build wasi || wasm
// +build wasi wasm

package messagingSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/smartops
//export getMessagingPubSubName
func GetName(resourceId uint32, data *byte) (error errno.Error)

//go:wasm-module taubyte/smartops
//export getMessagingPubSubNameSize
func GetNameSize(resourceId uint32, size *uint32) (error errno.Error)
