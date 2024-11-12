package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
)

type TestApi struct {
	v1 *v1.Router
}

func NewTestApi(v1 *v1.Router) *TestApi {
	return &TestApi{
		v1: v1,
	}
}
func (u *TestApi) Info(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong---2--test--Info",
	})
}
func (u *TestApi) Router() {
	var route = u.v1.GetNotLogin("test")
	route.GET("info", u.Info)
}
