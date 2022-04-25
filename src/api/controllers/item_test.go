package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mercadolibre/fury-core-go-template/src/api/domain/models"
	"github.com/stretchr/testify/assert"
)

func Test_CreateItem(t *testing.T) {

	body := NewItemBody()
	res := NewRequest(srv, http.MethodPost, "/item", body)

	resItem := models.Item{}
	_ = json.Unmarshal(res.Body.Bytes(), &resItem)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "aName", resItem.Name)
}

func Test_CreateItemBadRequest(t *testing.T) {

	res := NewRequest(srv, http.MethodPost, "/item", "bad")

	resError := errModel{}
	_ = json.Unmarshal(res.Body.Bytes(), &resError)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "invalid character 'b' looking for beginning of value", resError.Description)
}

func Test_GetItem(t *testing.T) {

	res := NewRequest(srv, http.MethodGet, "/item/1", "")

	resItem := models.Item{}
	_ = json.Unmarshal(res.Body.Bytes(), &resItem)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "aName", resItem.Name)
}

func Test_GetItemNotFound(t *testing.T) {

	body := NewItemBody()
	res := NewRequest(srv, http.MethodGet, "/item/42", body)

	resError := errModel{}
	_ = json.Unmarshal(res.Body.Bytes(), &resError)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "Item not found", resError.Error)
}

func NewItemBody() string {
	return `{
		"name": "aName"
	}`
}
