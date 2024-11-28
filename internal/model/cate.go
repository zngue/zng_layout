package model

import (
	"github.com/zngue/zng_app/db/data"
	"gorm.io/gorm"
)

type Cate struct {
	Id        string `gorm:"column:id" json:"id"`
	CateName  string `gorm:"column:cate_name" json:"cateName"`
	CateEName string `gorm:"column:cate_e_name" json:"cateEName"`
}

// TableName 表名
func (c *Cate) TableName() string {
	return "host_domain_cat"
}
func NewCate(conn *gorm.DB) *data.DB[Cate] {
	return data.NewDB[Cate](conn)
}
