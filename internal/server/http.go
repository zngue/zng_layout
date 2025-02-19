package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/pkg/types"
	"github.com/zngue/zng_layout/internal/conf"
	"net/http"
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
