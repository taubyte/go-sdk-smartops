package storage

import (
	"fmt"

	storageSym "github.com/taubyte/go-sdk-smartops/symbols/resource/storage"
)

func (f *Storage) Name() (string, error) {
	var nameSize uint32
	err := storageSym.GetNameSize(f.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting storage name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a storage with no name")
	}

	name := make([]byte, nameSize)
	err = storageSym.GetName(f.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting storage name failed with: %s", err)
	}

	return string(name), nil
}
