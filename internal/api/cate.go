package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/db/data"
	"github.com/zngue/zng_app/db/data/page"
	"github.com/zngue/zng_app/db/data/where"
	"github.com/zngue/zng_app/log"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
	"github.com/zngue/zng_layout/internal/model"
	"io"
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
		app.ApiFn(route, app.GET, "content", c.Content),
		app.ApiFn(route, app.GET, "list", c.List),
		app.ApiFn(route, app.POST, "transOut", c.TransOut),
		app.ApiFn(route, app.POST, "transIn", c.TransIn),
		app.ApiGetFn(route, "transOutCompensate", c.TransOutCompensate),
		app.ApiGetFn(route, "transInCompensate", c.TransInCompensate),
	)
}

//TransOutCompensate

func (c *CateApi) TransOutCompensate(ctx *gin.Context) (rs any, err error) {
	fmt.Println("TransOutCompensate")
	rs = api.NewDataApi(320, map[string]any{
		"dtm_result": "FAILURE",
		"message":    "current status 'failed', cannot prepare TransOutCompensate",
	})
	return
}
func (c *CateApi) TransInCompensate(ctx *gin.Context) (rs any, err error) {
	fmt.Println("TransInCompensate")
	err = api.NewError(409, map[string]any{
		"dtm_result": "FAILURE",
		"message":    "current status 'failed', cannot prepare TransOutCompensate",
	})
	return
}

func (c *CateApi) TransOut(ctx *gin.Context) (rs any, err error) {
	fmt.Println(io.ReadAll(ctx.Request.Body))
	fmt.Println("TransOut")
	err = errors.New("TransOut")
	return
}
func (c *CateApi) TransIn(ctx *gin.Context) (rs any, err error) {
	fmt.Println("TransIn")
	fmt.Println(io.ReadAll(ctx.Request.Body))
	err = errors.New("TransIn")
	return
}

func (c *CateApi) List(ctx *gin.Context) (rs any, err error) {
	list, err := c.cateConn.ListFn(
		data.PageWithData(
			page.DataWithPage(-1),
		),
		data.WhereOption(),
	)
	if err != nil {
		return
	}
	resData, err := c.cateConn.ContentFn(data.WhereOption(
		where.DataWhereOption("id", where.Gt, 100),
	))
	if err != nil {
		log.Errorf("test，%s", "刪除信息")
		err = api.ErrParameter
		return
	}
	rs = map[string]any{
		"content": resData,
		"list":    list,
	}
	return
}
func (c *CateApi) Content(ctx *gin.Context) (rs any, err error) {
	rs, err = c.cateConn.ContentFn(data.WhereOption(where.DataWhereOption("id", where.Like, "6")))
	return
}
