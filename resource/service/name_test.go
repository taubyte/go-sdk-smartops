package service_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	serviceSym "github.com/taubyte/go-sdk-smartops/symbols/resource/service"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunction(t *testing.T) {
	m := serviceSym.MockData{
		Id:   0,
		Name: "someService",
	}.Mock()

	var r resource.Resource
	msg, err := r.Service()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := msg.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someService" {
		t.Errorf("Expected name to be 'someService', got '%s'", name)
		return
	}

	// Size Method error
	m.Id = 1
	m.Mock()
	_, err = msg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	m.Id = 0
	m.Name = ""
	m.Mock()
	_, err = msg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data Method error
	m.Name = "someService"
	m.Mock()

	serviceSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = msg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
