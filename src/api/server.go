package api

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"github.com/singhkshitij/golang-rest-service-starter/src/metrics"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	AddMiddlewares(r)
	MakeHandler(r)

	address := ":" + config.Port()
	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return &Server{srv}, nil
}

func AddMiddlewares(router *gin.Engine) {

	//Adding CORS middleware
	var allowedOrigins []string
	if config.GetEnv() == "prod" {
		allowedOrigins = []string{"https://pro.superdms.app"}
	} else {
		allowedOrigins = []string{"http://localhost"}
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	logger.RegisterLoggerForRouter(router)
	router.Use(stats.RequestStats())

	metrics.GetMetricsMonitor().Use(router)
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
