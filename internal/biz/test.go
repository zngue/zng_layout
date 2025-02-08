package biz

import (
	"github.com/gin-gonic/gin"
)

type BizRepo interface {
}
type BizUseCase struct {
	repo BizRepo
}

func NewBizUseCase(repo BizRepo) *BizUseCase {
	return &BizUseCase{repo: repo}
}

func (b *BizUseCase) Abc(ctx *gin.Context, s string) (err error) {

	return
}
