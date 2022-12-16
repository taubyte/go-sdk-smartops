package service

import (
	"fmt"

	serviceSym "github.com/taubyte/go-sdk-smartops/symbols/resource/service"
)

func (f *Service) Name() (string, error) {
	var nameSize uint32
	err := serviceSym.GetNameSize(f.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting service name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a service with no name")
	}

	name := make([]byte, nameSize)
	err = serviceSym.GetName(f.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting service name failed with: %s", err)
	}

	return string(name), nil
}
