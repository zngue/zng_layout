package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/api"
	"io"
)

func NewHttp() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		api.DataSuccess(ctx)
		return
	})
	engine.NoRoute(func(ctx *gin.Context) {
		api.DataSuccess(ctx, api.Code(404), api.Msg("404"))
	})
	group := engine.Group("v1/dtm")
	{
		group.POST("transOut", func(context *gin.Context) {
			gid, _ := context.GetQuery("gid")
			var data, err = io.ReadAll(context.Request.Body)
			fmt.Println(gid, string(data), err)
			context.JSON(419, map[string]string{
				"dtm_result": "FAILURE",
				"message":    "current status 'failed', cannot prepare TransOut",
			})
		})
		group.POST("transOutCompensate", func(context *gin.Context) {
			gid, _ := context.GetQuery("gid")
			var data, err = io.ReadAll(context.Request.Body)
			fmt.Println(gid, string(data), err)
			context.JSON(409, map[string]string{
				"dtm_result": "FAILURE",
				"message":    "current status 'failed', cannot prepare TransOutCompensate",
			})
		})
		group.POST("transIn", func(context *gin.Context) {
			gid, _ := context.GetQuery("gid")
			var data, err = io.ReadAll(context.Request.Body)
			fmt.Println(gid, string(data), err)
			context.JSON(409, map[string]string{
				"dtm_result": "FAILURE",
				"message":    "current status 'failed', cannot prepare TransIn",
			})
		})
		group.POST("transInCompensate", func(context *gin.Context) {
			var data, err = io.ReadAll(context.Request.Body)
			gid, _ := context.GetQuery("gid")
			fmt.Println(gid, string(data), err)
			context.JSON(409, map[string]string{
				"dtm_result": "FAILURE",
				"message":    "current status 'failed', cannot prepare TransInCompensate",
			})
		})
	}
	return engine
}
func NewHttpGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("")
}
