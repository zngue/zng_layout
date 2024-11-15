package model

import (
	"github.com/zngue/zng_app/db/data"
	"gorm.io/gorm"
)

type Test struct {
}

func NewTest(conn *gorm.DB) *data.DB[Test] {
	return data.NewDB[Test](conn)
}
