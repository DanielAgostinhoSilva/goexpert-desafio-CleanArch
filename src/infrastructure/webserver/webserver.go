package webserver

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Method      string
	HandlerFunc http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = Handler{Method: method, HandlerFunc: handler}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.HandlerFunc)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
