package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) Create(reqBody *cartModel.CreateReqBody) (cartModel.Cart, error) {
	var carts cartModel.Cart

	tx := r.db.Begin()

	result := tx.Debug().Table("carts").Create(&cartModel.Cart{
		UserId: reqBody.UserId,
		Status: "active",
	})
	if result.Error != nil {
		tx.Rollback()
		return carts, result.Error
	}
	tx.Commit()
    return carts, nil
}