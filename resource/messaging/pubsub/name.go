package messaging

import (
	"fmt"

	messagingSym "github.com/taubyte/go-sdk-smartops/symbols/resource/messaging/pubsub"
)

func (f *Messaging) Name() (string, error) {
	var nameSize uint32
	err := messagingSym.GetNameSize(f.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting messaging name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a messaging with no name")
	}

	name := make([]byte, nameSize)
	err = messagingSym.GetName(f.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting messaging name failed with: %s", err)
	}

	return string(name), nil
}
