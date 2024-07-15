package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (r *repo) Create(product *productModel.Product) (err error) {
	return r.db.Debug().Create(product).Error
}
