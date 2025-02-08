package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/zngue/zng_app/db/api"
	"reflect"
)

func NewHttp() *gin.Engine {
	binding.Query.Name()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("json") // 改为使用 `json` 标签
		})
	}
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
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
