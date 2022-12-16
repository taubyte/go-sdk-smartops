package database

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Resource) (*Database, error) {
	err := r.IsType(common.ResourceTypeDatabase)
	if err != nil {
		return nil, err
	}

	return &Database{r}, nil
}
