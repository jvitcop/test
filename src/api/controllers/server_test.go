package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/services"
	"github.com/mercadolibre/fury-core-go-template/src/api/repositories"
	logger "github.com/sirupsen/logrus"
)

var srv = NewTestServer()

type TestServer struct {
	*Server
}

func NewTestServer() *TestServer {
	config.Test = true

	db := repositories.ConnectDb()
	itemRepo := repositories.NewItemRepository(db)
	items := services.NewItemService(itemRepo)

	core := &domain.Core{
		ItemService: items,
	}

	repositories.Ensure()
	level, _ := logger.ParseLevel("debug")
	logger.SetLevel(level)
	gin.SetMode(gin.TestMode)
	srv := NewServer(core, gin.New())
	srv.ConfigureRouter()
	return &TestServer{
		Server: srv,
	}
}

func NewRequest(s *TestServer, method string, url string, body string, headers ...[]string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	for _, header := range headers {
		req.Header.Set(header[0], header[1])
	}
	s.Router.ServeHTTP(res, req)
	return res
}

type errModel struct {
	Description string `json:"description"`
	Error       string `json:"error"`
}
