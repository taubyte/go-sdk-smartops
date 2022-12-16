//go:build wasi || wasm
// +build wasi wasm

package resourceSym

import (
	"github.com/taubyte/go-sdk-smartops/common"
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/smartops
//export getEventId
func GetEventId(resourceId uint32, eventIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/smartops
//export getResourceType
func GetResourceType(resourceId uint32, resourceTypePtr *common.ResourceType) (error errno.Error)
