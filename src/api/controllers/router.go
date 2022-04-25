package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/mercadolibre/fury-core-go-template/src/api/controllers/middlewares"
)

func (s *Server) ConfigureRouter() {
	s.Router.Use(gin.Recovery())
	s.Router.RedirectFixedPath = false
	s.Router.RedirectTrailingSlash = true

	if !config.Test {
		s.Router.Use(middlewares.NewRelic())
	}

	// Health Check
	s.Router.GET("/ping", Ping)

	s.Router.GET("/item/:id", s.GetItem)

	s.Router.POST("/item", s.CreateItem)

}
