package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/db/data"
	"github.com/zngue/zng_app/db/data/page"
	"github.com/zngue/zng_app/db/data/where"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
	"github.com/zngue/zng_layout/internal/model"
)

type CateApi struct {
	v1       *v1.Router
	cateConn *data.DB[model.Cate]
	app.ApiService
}

func NewCateApi(
	v1 *v1.Router,
	cateConn *data.DB[model.Cate],
) *CateApi {
	return &CateApi{
		v1:       v1,
		cateConn: cateConn,
	}
}
func (c *CateApi) Run() []*app.Api {
	route := c.v1.GetNotLogin("cate")
	return app.ApiServiceFn(
		app.ApiFn(route, app.GET, "list", c.Content),
	)
}

func (c *CateApi) List(ctx *gin.Context) {
	list, err := c.cateConn.ListFn(
		data.PageWithData(
			page.DataWithPage(-1),
		),
	)
	resData, err := c.cateConn.ContentFn(data.WhereOption(
		where.DataWhereOption("ids", where.Gt, 100),
	))
	if err != nil {
		api.DataError(ctx, err)
		return
	}
	api.DataWithErr(ctx, err, map[string]any{
		"list": list,
		"data": resData,
	})
}
func (c *CateApi) Content(ctx *gin.Context) (data any, err error) {

	return
}
