package auth

import (
	"github.com/AwahMarwah/Technical_Test_3/repository/user"
	"github.com/AwahMarwah/Technical_Test_3/service/auth"
	"gorm.io/gorm"
)

type controller struct {
	authService auth.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{authService: auth.NewService(user.NewRepo(db))}
}
