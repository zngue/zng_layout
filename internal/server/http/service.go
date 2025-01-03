package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_layout/internal/conf"
	"net/http"
)

func NewService(c *conf.Bootstrap, handler *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", c.App.Port),
		Handler: handler,
	}
}
