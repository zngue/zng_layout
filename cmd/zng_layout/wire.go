//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/internal/cron"
	"github.com/zngue/zng_layout/internal/model"
	"github.com/zngue/zng_layout/internal/server"
)

// initApp init zng_app application.
func initApp(cfg *conf.Bootstrap) (*app.App, func(), error) {
	panic(wire.Build(
		model.ProviderSet,
		api.ProviderSet,
		server.ProviderSet,
		cron.ProviderSet,
		app.ProviderSet,
	))

}
