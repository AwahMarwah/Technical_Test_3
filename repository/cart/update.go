package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) Update(id *uint32, values *map[string]any) (err error) {
	return r.db.Model(cartModel.Cart{Id: *id}).Updates(values).Error
}
