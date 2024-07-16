package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) Create(reqBody *cartModel.CreateReqBody) (cartModel.Cart, error) {
	var (
		carts cartModel.Cart
		cartItems []cartModel.CartItem
	)

	tx := r.db.Begin()

	result := tx.Table("carts").Create(&cartModel.Cart{
		UserId: reqBody.UserId,
		Status: "active",
	})
	if result.Error != nil {
		tx.Rollback()
		return carts, result.Error
	}
	// GET USER ID BEFORE COMMIT
	result.Scan(&carts)
	lengthCartItems := len(reqBody.CartItems)
	if lengthCartItems > 0 {
		for _, v := range reqBody.CartItems {
			cartItem := cartModel.CartItem{
				ProductId: v.ProductId,
				Quantity: v.Quantity,
				CartId: carts.Id,
			}
			cartItems = append(cartItems, cartItem)
		}
		resultCartItems := tx.Table("cart_items").Create(&cartItems)
		if resultCartItems.Error != nil {
			tx.Rollback()
			return carts, resultCartItems.Error
		}
	}
	tx.Commit()
    return carts, nil
}