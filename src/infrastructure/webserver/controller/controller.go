package controller

import "github.com/go-chi/chi/v5"

type Controller interface {
	Router(router chi.Router)
	Path() string
}
