package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/models"

	metrics "github.com/mercadolibre/fury-core-go-template/src/api/utils"
)

type ItemRepository struct {
	db *pg.DB
}

func NewItemRepository(db *pg.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Save(item *models.Item) error {
	defer metrics.RecordTime("db_query_time", []string{"entity:Item", "method:Save"})()

	return r.db.Insert(item)
}

func (r *ItemRepository) Get(id int64) (*models.Item, error) {
	defer metrics.RecordTime("db_query_time", []string{"entity:Item", "method:Get"})()

	var item models.Item

	err := r.db.Model(&item).Where("item.id = ?", id).Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return &item, err
	}
	return &item, err
}
