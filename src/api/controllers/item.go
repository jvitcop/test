package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	flerr "github.com/mercadolibre/fury-core-go-template/src/api/domain/errors"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/models"
)

func (s *Server) CreateItem(c *gin.Context) {
	var itemBody models.Item

	if err := c.ShouldBindJSON(&itemBody); err != nil {
		RespondWithHttpError(c, "Invalid item data", http.StatusBadRequest, err)
		return
	}

	err := s.Core.ItemService.CreateItem(&itemBody)
	if err != nil {
		var ierr *flerr.InputError
		httpStatus := http.StatusInternalServerError
		if errors.As(err, &ierr) {
			httpStatus = http.StatusBadRequest
		}
		RespondWithHttpError(c, "Something went wrong", httpStatus, err)
		return
	}

	c.JSON(http.StatusOK, itemBody)
}

func (s *Server) GetItem(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		RespondWithHttpError(c, "Could not retrieve item, id should be a number", http.StatusBadRequest, err)
		return
	}

	item, err := s.Core.ItemService.Repository().Get(int64(id))
	if err != nil {
		RespondWithHttpError(c, "Could not get item", http.StatusInternalServerError, err)
		return
	}

	if item == nil {
		RespondWithHttpError(c, "Item not found", http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, item)

}
