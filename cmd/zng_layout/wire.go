//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
	"github.com/zngue/zng_layout/internal/http"
)

// initApp init zng_app application.
func initApp(port int) (*app.App, func(), error) {
	panic(
		wire.Build(
			api.ProviderSet,
			http.ProviderSet,
			app.NewApp,
		),
	)

}
