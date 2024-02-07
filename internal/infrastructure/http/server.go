package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPHandler interface {
	Setup(e *echo.Echo)
}

type HTTPServer struct {
	Engine *echo.Echo
	port   string
}

func NewHTTPServer(port string, handlers ...HTTPHandler) *HTTPServer {
	server := &HTTPServer{
		port:   port,
		Engine: echo.New(),
	}

	server.Engine.Use(middleware.Logger())

	for _, h := range handlers {
		h.Setup(server.Engine)
	}

	return server
}

func (s *HTTPServer) Start() error {
	return s.Engine.Start(fmt.Sprintf(":%s", s.port))
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.Engine.Shutdown(ctx)
}
