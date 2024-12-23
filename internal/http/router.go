package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/log"
)

func NewHttp() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger(), log.RequestGinLog())
	engine.GET("/ping", func(ctx *gin.Context) {
		api.DataSuccess(ctx)
		return
	})
	engine.NoRoute(func(ctx *gin.Context) {
		api.DataSuccess(ctx, api.Code(404), api.Msg("404"))
		//return
	})
	return engine
}
func NewHttpGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("")
}
