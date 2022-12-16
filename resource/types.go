package resource

import (
	"fmt"

	"github.com/taubyte/go-sdk-smartops/common"
	resourceSym "github.com/taubyte/go-sdk-smartops/symbols/resource"
)

var _ common.Resource = Resource(0)

type Resource uint32

func (r Resource) Type() (common.ResourceType, error) {
	var resourceType common.ResourceType

	err := resourceSym.GetResourceType(uint32(r), &resourceType)
	if err != 0 {
		return 0, fmt.Errorf("Getting resource type failed with: %s", err)
	}

	return resourceType, nil
}

func (r Resource) IsType(expectedType common.ResourceType) error {
	resourceType, err := r.Type()
	if err != nil {
		return err
	}

	if resourceType != expectedType {
		return fmt.Errorf("Resource is not %s got %s", expectedType, resourceType)
	}

	return nil
}

func (r Resource) Uint32() uint32 {
	return uint32(r)
}
