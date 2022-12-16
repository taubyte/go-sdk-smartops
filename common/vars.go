package common

type ResourceType uint32

const (
	ResourceTypeUndefined ResourceType = iota
	ResourceTypeFunctionHTTP
	ResourceTypeWebsite
	ResourceTypeFunctionP2P
	ResourceTypeFunctionPubSub
	ResourceTypeDatabase
	ResourceTypeStorage
	ResourceTypeService
	ResourceTypeMessagingPubSub
	ResourceTypeMessagingWebSocket

	// Marks the end
	ResourceTypeCap
)

var resourceStrings = []string{
	"ResourceTypeUndefined",
	"ResourceTypeFunctionHTTP",
	"ResourceTypeWebsiteHTTP",
	"ResourceTypeFunctionP2P",
	"ResourceTypeFunctionPubSub",
	"ResourceTypeDatabase",
	"ResourceTypeStorage",
	"ResourceTypeService",
	"ResourceTypeMessagingPubSub",
	"ResourceTypeMessagingWebSocket",
}

func (et ResourceType) String() string {
	if et < ResourceTypeUndefined || et > ResourceTypeCap {
		return resourceStrings[0]
	}

	return resourceStrings[et]
}
