package storage

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Resource) (*Storage, error) {
	err := r.IsType(common.ResourceTypeStorage)
	if err != nil {
		return nil, err
	}

	return &Storage{r}, nil
}
