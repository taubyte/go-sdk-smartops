package resource

import (
	"github.com/taubyte/go-sdk-smartops/resource/database"
	"github.com/taubyte/go-sdk-smartops/resource/function"
	"github.com/taubyte/go-sdk-smartops/resource/messaging"
	"github.com/taubyte/go-sdk-smartops/resource/service"
	"github.com/taubyte/go-sdk-smartops/resource/storage"
	"github.com/taubyte/go-sdk-smartops/resource/website"
)

func (r Resource) Database() (*database.Database, error) {
	return database.Convert(r)
}

func (r Resource) Function() function.Function {
	return function.Convert(r)
}

func (r Resource) Messaging() messaging.Messaging {
	return messaging.Convert(r)
}

func (r Resource) Service() (*service.Service, error) {
	return service.Convert(r)
}

func (r Resource) Storage() (*storage.Storage, error) {
	return storage.Convert(r)
}

func (r Resource) Website() (*website.Website, error) {
	return website.Convert(r)
}
