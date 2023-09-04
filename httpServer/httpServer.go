package httpServer

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type httpServer struct {
	Server *echo.Echo
}

func NewHTTPServer() *httpServer {
	return &httpServer{
		Server: echo.New(),
	}
}

func (s *httpServer) Start() {
	log.Printf("server is running on port %s", "8080")
	go func() {
		if err := s.Server.Start(":8080"); err != nil {
			log.Printf("server start failed: %v", err)
		}
	}()
}

func (s *httpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %v", err)
	}
	log.Print("Gracefully shutdown the server")
}
