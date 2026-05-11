package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	dotenv "github.com/joho/godotenv"
	"github.com/open-portfolios/review/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf config.yaml")
	if err := dotenv.Load(); err != nil {
		log.Info(err)
	}
}

func main() {
	c := config.New(config.WithSource(
		env.NewSource("REVIEW_"),
		file.NewSource(flagconf),
	))
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(bc.Data.Database.Source))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/data/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	defer g.Execute()

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
}
