package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (r *repo) Find(selectParams []string, conditions *productModel.Product) (resData []productModel.Product, err error) {
	return resData, r.db.Select(selectParams).Find(&resData, conditions).Error
}
