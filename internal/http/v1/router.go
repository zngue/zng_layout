package v1

import "github.com/gin-gonic/gin"

type Router struct {
	router      *gin.RouterGroup //根据情况自己处理
	loginRouter *gin.RouterGroup //需要登录的路由
}

func (r *Router) GetNotLogin(names ...string) (route *gin.RouterGroup) {
	if len(names) > 0 {
		return r.router.Group(names[0])
	}
	return r.router
}
func (r *Router) GetLogin(names ...string) (route *gin.RouterGroup) {
	if len(names) > 0 {
		return r.loginRouter.Group(names[0])
	}
	return r.loginRouter
}

func NewRouter(group *gin.RouterGroup) *Router {
	v1 := group.Group("v1")
	v1Login := group.Group("v1")
	v1Login.Use(func(context *gin.Context) {
		context.Set("userId", 1)
		return
	})
	return &Router{
		router:      v1,
		loginRouter: v1Login,
	}
}
