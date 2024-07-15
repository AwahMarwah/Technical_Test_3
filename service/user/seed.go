package user

import (
	"errors"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/encrypt"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"gorm.io/gorm"
)

func (s *service) Seed(req *userModel.SeedReq) (err error) {
	if _, err = s.userRepo.Take([]string{"id"}, &userModel.User{Email: req.Email}); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	} else {
		return errors.New(common.EmailAlreadyExists)
	}
	if err = encrypt.GenerateFromPassword(&req.Password); err != nil {
		return
	}
	user := userModel.User{
		Email: req.Email,
		EncryptedPassword: req.Password,
	}
	return s.userRepo.Create(&user)
}