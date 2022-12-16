package function_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	resourceSym "github.com/taubyte/go-sdk-smartops/symbols/resource"
)

func TestFunctionNewError(t *testing.T) {
	var r resource.Resource
	resourceSym.MockType(0, 0)

	_, err := r.Function().Http()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
