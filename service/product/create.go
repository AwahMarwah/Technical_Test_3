package product

import (
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
)

func (s *service) Create( reqBody *productModel.CreateReqBody) (err error) {
	return s.productRepo.Create(&productModel.Product{
		Name: reqBody.Name,
		Price: reqBody.Price,
		Description: reqBody.Description,
	})
}