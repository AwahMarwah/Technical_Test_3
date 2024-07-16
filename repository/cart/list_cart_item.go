package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) ListCartItem(id uint32) (resData []cartModel.CartItem, err error) {
	resData = make([]cartModel.CartItem, 0)
	return resData, r.db.Select("id, product_id, quantity").Model(&cartModel.CartItem{}).Scan(&resData).Error
}
