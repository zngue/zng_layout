package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/db/data"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
	"github.com/zngue/zng_layout/internal/model"
)

type TestApi struct {
	v1         *v1.Router
	testConn   *data.DB[model.Test]
	userConn   *data.DB[model.User]
	memberConn *data.DB[model.Member]
}

func NewTestApi(
	v1 *v1.Router,
	testConn *data.DB[model.Test],
	userConn *data.DB[model.User],
	memberConn *data.DB[model.Member],
) *TestApi {
	return &TestApi{
		v1:         v1,
		testConn:   testConn,
		userConn:   userConn,
		memberConn: memberConn,
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
