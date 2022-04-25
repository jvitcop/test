package repositories

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/mercadolibre/fury-core-go-template/src/api/domain/models"
	logger "github.com/sirupsen/logrus"
)

func ConnectDb() *pg.DB {
	host := config.DbHostname
	if config.Test {
		host = "localhost"
	}
	db := pg.Connect(&pg.Options{
		User:     config.DbUsername,
		Password: config.DbPassword,
		Database: config.DbDatabase,
		Addr:     fmt.Sprintf("%s:%v", host, config.DbPort),
	})
	//todo settings

	db.AddQueryHook(dbLogger{})

	return db
}

func Ensure() {
	db := ConnectDb()

	for _, model := range []interface{}{
		(*models.Item)(nil),
	} {
		if config.Test {
			// Drop tables
			_ = db.DropTable(model, &orm.DropTableOptions{
				IfExists: true,
			})
		}
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			//Temp: config.Test,
		})
		if err != nil {
			panic(err)
		}
	}
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	if bs, err := q.FormattedQuery(); err == nil {
		logger.Trace(string(bs))
	}

	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}
