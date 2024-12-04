package http

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
)

func NewApiService(testApi *api.TestApi) []app.IApiService {
	return []app.IApiService{
		testApi,
	}
}
