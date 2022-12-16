//go:build !wasi && !wasm
// +build !wasi,!wasm

package resourceSym

import (
	"github.com/taubyte/go-sdk-smartops/common"
	"github.com/taubyte/go-sdk/errno"
)

var GetEventId = func(resourceId uint32, eventIdPtr *uint32) (error errno.Error) {
	return 0
}

var GetResourceType = func(resourceId uint32, typePtr *common.ResourceType) (error errno.Error) {
	return 0
}
