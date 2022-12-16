package messaging_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	messagingSym "github.com/taubyte/go-sdk-smartops/symbols/resource/messaging/websocket"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunction(t *testing.T) {
	m := messagingSym.MockData{
		Id:   0,
		Name: "someMessaging",
	}.Mock()

	var r resource.Resource
	msg, err := r.Messaging().WebSocket()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := msg.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someMessaging" {
		t.Errorf("Expected name to be 'someMessaging', got '%s'", name)
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
	m.Name = "someMessaging"
	m.Mock()

	messagingSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = msg.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
