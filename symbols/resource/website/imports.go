//go:build !wasi && !wasm
// +build !wasi,!wasm

package websiteSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var GetName = func(resourceId uint32, data *byte) (error errno.Error) {
	return 0
}

var GetNameSize = func(resourceId uint32, size *uint32) (error errno.Error) {
	return 0
}
