package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) Take(selectParams []string, conditions *cartModel.Cart) (cart cartModel.Cart, err error) {
	return cart, r.db.Select(selectParams).Take(&cart, conditions).Error
}
