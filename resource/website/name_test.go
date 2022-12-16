package website_test

import (
	"testing"

	"github.com/taubyte/go-sdk-smartops/resource"
	websiteSym "github.com/taubyte/go-sdk-smartops/symbols/resource/website"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunction(t *testing.T) {
	m := websiteSym.MockData{
		Id:   0,
		Name: "someWebsite",
	}.Mock()

	var r resource.Resource
	website, err := r.Website()
	if err != nil {
		t.Error(err)
		return
	}

	name, err := website.Name()
	if err != nil {
		t.Error(err)
		return
	}

	if name != "someWebsite" {
		t.Errorf("Expected name to be 'someWebsite', got '%s'", name)
		return
	}

	// Size Method error
	m.Id = 1
	m.Mock()
	_, err = website.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Size 0
	m.Id = 0
	m.Name = ""
	m.Mock()
	_, err = website.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}

	// Data Method error
	m.Name = "someWebsite"
	m.Mock()

	websiteSym.GetName = func(resourceId uint32, data *byte) (error errno.Error) {
		return 1
	}

	_, err = website.Name()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
