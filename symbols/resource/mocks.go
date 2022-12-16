//go:build !wasi && !wasm
// +build !wasi,!wasm

package resourceSym

import (
	"github.com/taubyte/go-sdk-smartops/common"
	"github.com/taubyte/go-sdk/errno"
)

func MockType(testId uint32, testType common.ResourceType) {
	GetResourceType = func(resourceId uint32, typePtr *common.ResourceType) (error errno.Error) {
		if testId != resourceId {
			return 1
		}

		*typePtr = testType

		return 0
	}
}
