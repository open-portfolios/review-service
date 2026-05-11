//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/open-portfolios/review/internal/biz"
	"github.com/open-portfolios/review/internal/conf"
	"github.com/open-portfolios/review/internal/data"
	"github.com/open-portfolios/review/internal/infra"
	"github.com/open-portfolios/review/internal/server"
	"github.com/open-portfolios/review/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Snowflake, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		infra.ProviderSet,
		newApp,
	))
}
