package storage_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	storageSym "github.com/taubyte/go-sdk-smartops/symbols/resource/storage"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunction(t *testing.T) {
	m := storageSym.MockData{
		Id:   0,
		Name: "someStorage",
	}.Mock()

	var r resource.Resource
	stg, err := r.Storage()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := stg.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someStorage" {
		t.Errorf("Expected name to be 'someStorage', got '%s'", name)
		return
	}

	// Size Method error
	m.Id = 1
	m.Mock()
	_, err = stg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	m.Id = 0
	m.Name = ""
	m.Mock()
	_, err = stg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data Method error
	m.Name = "someStorage"
	m.Mock()

	storageSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = stg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
