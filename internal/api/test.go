package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/db/data"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
	"github.com/zngue/zng_layout/internal/model"
)

type TestApi struct {
	v1       *v1.Router
	testConn *data.DB[model.Test]
	app.ApiService
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
func (u *TestApi) Run() []*app.Api {
	route := u.v1.GetNotLogin("test")
	return app.ApiServiceFn(
		app.ApiFn(route, app.GET, "list", u.Content),
	)
}
