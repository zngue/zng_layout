package db

type User struct {
	Id        int32  `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Phone     string `gorm:"column:phone"`
	Status    int32  `gorm:"column:status"`
	CreatedAt int32  `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
	Sex       int32  `gorm:"column:sex"`
	Avatar    string `gorm:"column:avatar"`
}

func (User) TableName() string
func (User) TableName() string {
	return "user"
}
