package user

import (
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
	"github.com/AwahMarwah/Technical_Test_3/service/user"
	"gorm.io/gorm"
)

type controller struct {
	userService user.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{userService: user.NewService(userRepo.NewRepo(db))}
}
