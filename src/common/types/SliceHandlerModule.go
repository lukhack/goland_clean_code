package common

import "github.com/gofiber/fiber/v3"

type HandlerModule struct {
	Route   string
	Method  interface{}
	Handler func(ctx fiber.Ctx) error
}

type SliceHandlers struct {
	Prefix string
	Routes []HandlerModule
}

type GlobalHandler []SliceHandlers

type HandlersStore struct {
	Handlers []SliceHandlers
}

func NewHandlersStore() *HandlersStore {
	return &HandlersStore{}
}
