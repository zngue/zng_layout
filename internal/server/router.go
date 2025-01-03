package server

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_layout/internal/api"
)

type RouterApi []router.IApiService

//合并路由 CombineRouter

func NewCombineRouter(routerItems RouterApi) (routes []router.IApiService) {
	routes = append(routes, routerItems...)
	return
}
func NewRouter(testApi *api.TestApi) (routerItems RouterApi) {
	routerItems = []router.IApiService{
		testApi,
	}
	return
}

var ProviderSetRouter = wire.NewSet(
	NewRouter,
	NewCombineRouter,
)
