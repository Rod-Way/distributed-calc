package http

import (
	"context"
	"distributedCalculator/internal/server/http/handlers"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	EchoServer *echo.Echo
}

func New() *Server {
	e := echo.New()
	return &Server{EchoServer: e}
}

func (s *Server) Start() {
	s.applyHandlers()
	s.EchoServer.Logger.Fatal(s.EchoServer.Start(":1323"))
}

func (s *Server) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.EchoServer.Shutdown(ctx); err != nil {
		s.EchoServer.Logger.Fatal(err)
	}
}

// func (s *Server) applyMiddlewares() {}

func (s *Server) applyHandlers() {
	h := handlers.NewHandler()
	s.EchoServer.POST("api/v1/calculate", h.PostCalculate)
}
