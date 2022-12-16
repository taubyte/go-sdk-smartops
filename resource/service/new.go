package service

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Resource) (*Service, error) {
	err := r.IsType(common.ResourceTypeService)
	if err != nil {
		return nil, err
	}

	return &Service{r}, nil
}
