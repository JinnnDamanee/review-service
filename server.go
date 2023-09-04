package main

import (
	"jindamanee2544/review-service/httpServer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Review struct {
	ReviewID   int    `json:"review_id"`
	ReviewName string `json:"review_name"`
}

func main() {
	s := httpServer.NewHTTPServer()

	s.Server.GET("/review", func(c echo.Context) error {
		log.Print("GET /review")
		r := Review{
			ReviewID:   1,
			ReviewName: "review name",
		}

		return c.JSON(http.StatusOK, r)
	})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.Start()
	}()

	<-sig

	s.Shutdown()
}
