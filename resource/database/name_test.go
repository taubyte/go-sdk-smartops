package database_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	databaseSym "github.com/taubyte/go-sdk-smartops/symbols/resource/database"
	"github.com/taubyte/go-sdk/errno"
)

func TestDatabase(t *testing.T) {
	m := databaseSym.MockData{
		Id:   0,
		Name: "someDatabase",
	}.Mock()

	var r resource.Resource
	db, err := r.Database()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := db.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someDatabase" {
		t.Errorf("Expected name to be 'someDatabase', got '%s'", name)
		return
	}

	// Size Method error
	m.Id = 1
	m.Mock()
	_, err = db.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	m.Id = 0
	m.Name = ""
	m.Mock()
	_, err = db.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data Method error
	m.Name = "someDatabase"
	m.Mock()

	databaseSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = db.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
