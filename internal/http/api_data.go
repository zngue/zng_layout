package http

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
)

func NewApiService(testApi *api.TestApi, cateApi *api.CateApi) []app.IApiService {
	return []app.IApiService{
		testApi,
		cateApi,
	}
}
