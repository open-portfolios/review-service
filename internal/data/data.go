package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/open-portfolios/review/internal/conf"
	"github.com/open-portfolios/review/internal/data/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewReviewRepo)

type Data struct {
	q *query.Query
}

func NewData(db *gorm.DB, logger log.Logger) (*Data, func()) {
	query.SetDefault(db)
	d := &Data{q: query.Q}
	cleanup := func() {
		log.NewHelper(logger).Debug("cleaning up resources")
	}
	return d, cleanup
}

func NewDB(data *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(data.Database.Source))
	if err != nil {
		return nil, err
	}
	return db, nil
}
