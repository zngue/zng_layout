package model

import (
	"github.com/zngue/zng_app/db/data"
	"gorm.io/gorm"
)

type User struct {
}

func NewUser(conn *gorm.DB) *data.DB[User] {
	return data.NewDB[User](conn)
}
