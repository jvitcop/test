package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain"
	logger "github.com/sirupsen/logrus"
)

type Server struct {
	Core   *domain.Core
	Router *gin.Engine
}

func NewServer(core *domain.Core, router *gin.Engine) *Server {
	level, _ := logger.ParseLevel(config.LoggerLevel)
	logger.SetLevel(level)

	return &Server{Core: core, Router: router}
}
