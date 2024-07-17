package cart

import (
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"github.com/AwahMarwah/Technical_Test_3/repository/cart"
	productRepo "github.com/AwahMarwah/Technical_Test_3/repository/product"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
)

type (
	IService interface {
		Create(reqBody *cartModel.CreateReqBody) (err error)
		Detail(reqPath *cartModel.ReqPath) (resData cartModel.DetailResData, statusCode int, err error)
		List(reqQuery *cartModel.ListReqQuery) (resData []cartModel.ListCarts, count int64, err error)
		Update(reqPath *cartModel.UpdateReqPath) (statusCode int, err error)
	}

	service struct {
		cartRepo    cart.IRepo
		userRepo    userRepo.IRepo
		productRepo productRepo.IRepo
	}
)

func NewService(cartRepo cart.IRepo, userRepo userRepo.IRepo, productRepo productRepo.IRepo) IService {
	return &service{
		cartRepo:    cartRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}
