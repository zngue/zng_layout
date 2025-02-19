package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/pkg/types"
	"github.com/zngue/zng_layout/internal/conf"
	"net/http"
	"reflect"
)

func NewHttpService(c *conf.Bootstrap, handler *gin.Engine, routes []types.Register) *http.Server {
	for _, register := range routes {
		register.Register()
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", c.App.Port),
		Handler: handler,
	}
}
func NewHttpEngine() *gin.Engine {
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
	})
	return engine
}
func NewHttpGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("")
}
