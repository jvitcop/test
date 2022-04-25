package domain

import (
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/services"
	"github.com/mercadolibre/fury-core-go-template/src/api/repositories"
)

type Core struct {
	ItemService *services.ItemService
}

func NewCore() *Core {
	db := repositories.ConnectDb()

	repo := repositories.NewItemRepository(db)
	service := services.NewItemService(repo)

	return &Core{
		ItemService: service,
	}
}
