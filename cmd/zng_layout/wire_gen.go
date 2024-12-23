// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/internal/http"
	"github.com/zngue/zng_layout/internal/http/v1"
	"github.com/zngue/zng_layout/internal/model"
)

// Injectors from wire.go:

// initApp init zng_app application.
func initApp(bootstrap *conf.Bootstrap) (*app.App, func(), error) {
	engine := http.NewHttp()
	server := http.NewService(bootstrap, engine)
	routerGroup := http.NewHttpGroup(engine)
	router := v1.NewRouter(routerGroup)
	db, err := model.NewDB(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	dataDB := model.NewTest(db)
	testApi := api.NewTestApi(router, dataDB)
	v := http.NewApiService(testApi)
	v2 := app.NewRouter(v)
	v3, err := http.NewCron()
	if err != nil {
		return nil, nil, err
	}
	appApp := app.NewApp(server, v2, v3)
	return appApp, func() {
	}, nil
}
