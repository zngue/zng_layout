package model

import (
	"github.com/zngue/zng_app/db/data"
	"gorm.io/gorm"
)

type Member struct {
}

func NewMember(conn *gorm.DB) *data.DB[Member] {
	return data.NewDB[Member](conn)
}
