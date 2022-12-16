package website

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Resource) (*Website, error) {
	err := r.IsType(common.ResourceTypeWebsite)
	if err != nil {
		return nil, err
	}

	return &Website{r}, nil
}
