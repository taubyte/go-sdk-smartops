package function_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	functionSym "github.com/taubyte/go-sdk-smartops/symbols/resource/function/pubsub"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunction(t *testing.T) {
	m := functionSym.MockData{
		Id:   0,
		Name: "someFunction",
	}.Mock()

	var r resource.Resource
	_func, err := r.Function().PubSub()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := _func.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someFunction" {
		t.Errorf("Expected name to be 'someFunction', got '%s'", name)
		return
	}

	// Size Method error
	m.Id = 1
	m.Mock()
	_, err = _func.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	m.Id = 0
	m.Name = ""
	m.Mock()
	_, err = _func.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data Method error
	m.Name = "someFunction"
	m.Mock()

	functionSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = _func.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
