package v1

import (
	"github.com/gin-gonic/gin"
)

type Router gin.RouterGroup

func NewV1Router(api *gin.RouterGroup) *Router {
	return (*Router)(api.Group("v1"))
}
func (v *Router) Routes(name string) *gin.RouterGroup {
	api := (*gin.RouterGroup)(v)
	return api.Group(name)
}

type LoginV1Router gin.RouterGroup

func NewLoginV1Router(api *gin.RouterGroup) *LoginV1Router {
	return (*LoginV1Router)(api.Group("v1"))
}
func (v *LoginV1Router) Routes(name string) *gin.RouterGroup {
	api := (*gin.RouterGroup)(v)
	return api.Group(name)
}
