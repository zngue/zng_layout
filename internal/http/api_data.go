package http

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_layout/internal/api"
)

func NewRouter(testApi *api.TestApi) (routes []app.Router) {
	return []app.Router{
		testApi,
	}
}
