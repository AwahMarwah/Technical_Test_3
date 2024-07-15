package product

import (
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	productRepo "github.com/AwahMarwah/Technical_Test_3/repository/product"
)

type (
	IService interface {
		Create(reqBody *productModel.CreateReqBody) (err error)
	}

	service struct {
		productRepo productRepo.IRepo
	}
)

func NewService(productRepo productRepo.IRepo) IService {
	return &service{productRepo: productRepo}
}