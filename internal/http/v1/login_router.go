package v1

import (
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
	Router *gin.RouterGroup
}

func NewLoginRouter(api *gin.RouterGroup) *LoginRouter {
	return &LoginRouter{Router: api}
}
