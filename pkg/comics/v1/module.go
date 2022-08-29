package v1

import "sync"

type ComicModule struct {
	controller *ComicController
	service    *ComicService
}

var comicModuleSingleton *ComicModule
var comicModuleOnce sync.Once

func NewComicModule() *ComicModule {
	comicModuleOnce.Do(func() {
		service := NewComicService()
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
