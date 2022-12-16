package messaging

import (
	"github.com/taubyte/go-sdk-smartops/common"
)

func Convert(r common.Resource) (*Messaging, error) {
	err := r.IsType(common.ResourceTypeMessagingWebSocket)
	if err != nil {
		return nil, err
	}

	return &Messaging{r}, nil
}
