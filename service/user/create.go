package user

import (
	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/encrypt"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
)

func (s *service) Create(reqBody *userModel.CreateReq) (err error) {
	// CREATE DEFAULT PASSWORD
	if reqBody.EncryptedPassword == "" {
		reqBody.EncryptedPassword = common.DefaultPassword
	}
	// ENCRYPT PASS
	encPass, _ := encrypt.HashAndSalt([]byte(reqBody.EncryptedPassword))
	// ASSIGNE REQBODY TO MODEL USER
	user := userModel.User{
		Email:             reqBody.Email,
		EncryptedPassword: encPass,
	}
	return s.userRepo.Create(&user)
}
