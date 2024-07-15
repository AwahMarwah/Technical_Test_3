package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (r *repo) Take(selectParams []string, conditions *productModel.Product) (product productModel.Product, err error) {
	return product, r.db.Select(selectParams).Take(&product, conditions).Error
}
