package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/api"
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
	return engine
}
func NewHttpGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("")
}
