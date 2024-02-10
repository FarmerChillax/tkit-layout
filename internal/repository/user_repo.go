package repository

import (
	"github.com/FarmerChillax/tkit-layout/internal/model/mysql"
	"gorm.io/gorm"
)

type UserRepo struct {
	CommonRepository[mysql.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		CommonRepository: CommonRepository[mysql.User]{
			db: db.Table(mysql.User{}.TableName()),
		},
	}
}
