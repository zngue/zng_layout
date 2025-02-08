package server

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_layout/api/user/v1"
)

var ProviderSetRouter = wire.NewSet(
	NewRouter,
	NewCombineRouter,
	v1.NewUserGinHttpRouterService,
)
var ProviderSetPb = wire.NewSet(
	v1.NewUserGinHttpRouterService,
)

func NewRouter(testApi *v1.UserGinHttpRouterService) (routerItems RouterApi) {
	routerItems = []router.IApiService{
		testApi,
	}
	return
}

type RouterApi []router.IApiService

// NewCombineRouter 合并路由 CombineRouter
func NewCombineRouter(routerItems RouterApi) (routes []router.IApiService) {
	routes = append(routes, routerItems...)
	return
}
