package function

import (
	"fmt"

	functionSym "github.com/taubyte/go-sdk-smartops/symbols/resource/function/http"
)

func (f *Function) Name() (string, error) {
	var nameSize uint32
	err := functionSym.GetNameSize(f.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting function name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a function with no name")
	}

	name := make([]byte, nameSize)
	err = functionSym.GetName(f.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting function name failed with: %s", err)
	}

	return string(name), nil
}
