package http

import (
	"github.com/gin-gonic/gin"
)

func NewHttp() *gin.Engine {
	engine := gin.Default()
	return engine
}
func NewHttpGroup(engine *gin.Engine) *gin.RouterGroup {

	return engine.Group("")
}
