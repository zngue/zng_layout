//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/internal/http"
	"github.com/zngue/zng_layout/internal/model"
)

// initApp init zng_app application.
func initApp(*conf.Bootstrap) (*app.App, func(), error) {
	panic(wire.Build(
		model.ProviderSet,
		api.ProviderSet,
		http.ProviderSet,
		app.NewApp,
	))

}
