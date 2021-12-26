package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"net/http"
	"os"
	"os/signal"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() (*Server, error) {
	if config.GetEnv() == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	MakeHandler(r)

	address := ":" + config.Port()
	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return &Server{srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	logger.Debug("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	logger.Info("Listening on " + srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logger.Info("Shutting down server...", logger.KV("REASON", sig))
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	logger.Debug("Server gracefully stopped")
}
