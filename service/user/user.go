package user

import (
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
)

type (
	IService interface {
		Seed(req *userModel.SeedReq) (err error)
	}

	service struct {
		userRepo userRepo.IRepo
	}
)

func NewService(userRepo userRepo.IRepo) IService {
	return &service{userRepo: userRepo}
}