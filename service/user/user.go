package user

import (
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
)

type (
	IService interface {
		Create(reqBody *userModel.CreateReq) (err error)
		Detail(reqPath *userModel.ReqPath) (resData userModel.DetailUserResData, statusCod int, err error)
		List(reqQuery *userModel.ListReqQuery) (resData []userModel.ListUser, count int64, err error)
		Seed(req *userModel.SeedReq) (err error)
	}

	service struct {
		userRepo userRepo.IRepo
	}
)

func NewService(userRepo userRepo.IRepo) IService {
	return &service{userRepo: userRepo}
}
