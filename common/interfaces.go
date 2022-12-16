package common

type Resource interface {
	IsType(ResourceType) error
	Uint32() uint32
}

type Function interface {
	Resource
}
