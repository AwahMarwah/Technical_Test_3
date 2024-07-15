package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (s *service) List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListResData, count int64, err error) {
	return s.productRepo.List(reqQuery)
}