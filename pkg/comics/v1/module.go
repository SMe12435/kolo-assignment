package v1

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

type ComicModule struct {
	controller *ComicController
	service    *ComicService
}

var comicModuleSingleton *ComicModule
var comicModuleOnce sync.Once

func NewComicModule(redis *redis.Client) *ComicModule {
	comicModuleOnce.Do(func() {
		service := NewComicService(redis)
		controller := NewComicController()
		comicModuleSingleton = &ComicModule{
			service:    service,
			controller: controller,
		}
	})
	return comicModuleSingleton
}

func (m *ComicModule) GetController() *ComicController {
	return m.controller
}
