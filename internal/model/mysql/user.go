package mysql

import "gorm.io/plugin/soft_delete"

type User struct {
	Username string                `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password string                `gorm:"type:varchar(20);not null" json:"password"`
	IsDel    soft_delete.DeletedAt `json:"is_del" db:"is_del" gorm:"softDelete:flag,DeletedAtField:DeletedTime"`
}

func (User) TableName() string {
	return "user"
}
