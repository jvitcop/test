package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/mercadolibre/fury-core-go-template/src/api/controllers"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain"
	"github.com/mercadolibre/fury-core-go-template/src/api/repositories"
	logger "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	metrics "github.com/mercadolibre/fury-core-go-template/src/api/utils"
)

// @title Fury Core Go Template
// @version 0.0.1
// @description A template for our Go APIs
// @BasePath /
func main() {
	repositories.Ensure()
	run()
}

func run() {
	level, _ := logger.ParseLevel(config.LoggerLevel)
	logger.SetLevel(level)
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	core := domain.NewCore()

	metrics.ConnectDatadog()

	// Create web server
	srv := controllers.NewServer(core, gin.New())
	srv.ConfigureRouter()
	webServer := &http.Server{
		Addr:    ":8080",
		Handler: srv.Router,
	}

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			logger.WithError(err).Error("Server Shutdown")
		}
	}()

	// Start web server
	logger.Info("Starting Server")
	err := webServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.WithError(err).Error("Web Server Error")
	}

}
