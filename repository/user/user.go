package user

import (
	"gorm.io/gorm"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
)

type (
	IRepo interface {
		Create(user *userModel.User) (err error)
		Take(selectParams []string, conditions *userModel.User) (user userModel.User, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
