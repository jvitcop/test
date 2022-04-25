package services

import (
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/models"
)

type ItemReadRepository interface {
	Get(id int64) (*models.Item, error)
}
type ItemRepository interface {
	Save(item *models.Item) error
	ItemReadRepository
}

type ItemService struct {
	repo ItemRepository
}

func NewItemService(repo ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) Repository() ItemRepository {
	return s.repo
}

func (s *ItemService) CreateItem(item *models.Item) error {

	return s.repo.Save(item)

}
