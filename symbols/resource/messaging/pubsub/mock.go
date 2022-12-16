//go:build !wasi && !wasm
// +build !wasi,!wasm

package messagingSym

import (
	"github.com/taubyte/go-sdk-smartops/common"
	resourceSym "github.com/taubyte/go-sdk-smartops/symbols/resource"
	"github.com/taubyte/go-sdk-symbols/mocks"
)

type MockData struct {
	Id   uint32
	Name string
}

func (_m MockData) Mock() *MockData {
	m := &_m

	m.mockName()
	resourceSym.MockType(m.Id, common.ResourceTypeMessagingPubSub)

	return m
}

func (m *MockData) mockName() {
	GetName = mocks.DataIdStringFunction(m.Id, m.Name)
	GetNameSize = mocks.SizeIdStringFunction(m.Id, m.Name)
}
