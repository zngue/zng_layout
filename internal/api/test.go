package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/db/data"
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_layout/internal/model"
	v1 "github.com/zngue/zng_layout/internal/server/http/v1"
)

type TestApi struct {
	v1       *v1.Router
	testConn *data.DB[model.Test]
	router.ApiService
}

func NewTestApi(
	v1 *v1.Router,
	testConn *data.DB[model.Test],
) *TestApi {
	return &TestApi{
		v1:       v1,
		testConn: testConn,
	}
}

// Err
func (u *TestApi) Err(ctx *gin.Context) {

	api.DataError(ctx, fmt.Errorf("test"))
}
func (u *TestApi) Content(ctx *gin.Context) (data any, err error) {
	return
}
func (u *TestApi) Register() []*router.Api {
	route := u.v1.Routes("test")
	return router.ApiServiceFn(
		router.ApiFn(route, router.GET, "list", u.Content),
	)
}
