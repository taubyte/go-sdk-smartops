package website

import (
	"fmt"

	websiteSym "github.com/taubyte/go-sdk-smartops/symbols/resource/website"
)

func (f *Website) Name() (string, error) {
	var nameSize uint32
	err := websiteSym.GetNameSize(f.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting website name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a website with no name")
	}

	name := make([]byte, nameSize)
	err = websiteSym.GetName(f.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting website name failed with: %s", err)
	}

	return string(name), nil
}
