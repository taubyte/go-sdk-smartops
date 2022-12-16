package function

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Function) (*Function, error) {
	err := r.IsType(common.ResourceTypeFunctionP2P)
	if err != nil {
		return nil, err
	}

	return &Function{r}, nil
}
