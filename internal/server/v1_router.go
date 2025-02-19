package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zngue/zng_layout/api/user/v1"
	"github.com/zngue/zng_layout/internal/api"
)

func NewV1Router(router *gin.RouterGroup, userService *api.UserService) V1Router {
	return V1Router{
		v1.RegisterUserGinServer(router, userService),
	}
}
