package resource

import (
	"fmt"

	resourceSym "github.com/taubyte/go-sdk-smartops/symbols/resource"
	"github.com/taubyte/go-sdk/event"
	"github.com/taubyte/go-sdk/self"
)

func (r Resource) Project() (string, error) {
	return self.Project()
}

func (r Resource) Application() (string, error) {
	return self.Application()
}

func (r Resource) Event() (event.Event, error) {
	var eventId uint32
	err := resourceSym.GetEventId(uint32(r), &eventId)
	if err != 0 {
		return 0, fmt.Errorf("Getting event id failed with: %s", err)
	}

	return event.Event(eventId), nil
}
