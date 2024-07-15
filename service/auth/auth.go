package auth

import (
	authModel "github.com/AwahMarwah/Technical_Test_3/model/auth"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
)

type (
	IService interface {
		Authorize(token *string) (userId uint32, statusCode int, err error)
		SignIn(reqBody *authModel.SignInReqBody) (resData authModel.SignInResData, statusCode int, err error)
		SignOut(userId uint32) (err error)
	}

	service struct {
		userRepo userRepo.IRepo
	}
)

func NewService(userRepo userRepo.IRepo) IService {
	return &service{userRepo: userRepo}
}