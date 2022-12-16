package database_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	resourceSym "github.com/taubyte/go-sdk-smartops/symbols/resource"
)

func TestDatabaseNewError(t *testing.T) {
	var r resource.Resource
	resourceSym.MockType(0, 0)

	_, err := r.Database()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
