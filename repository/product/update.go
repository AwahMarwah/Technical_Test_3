package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (r *repo) Update(id *uint32, values *map[string]any) (err error) {
	return r.db.Model(productModel.Product{Id: *id}).Updates(values).Error
}
