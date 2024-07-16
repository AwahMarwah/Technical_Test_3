package cart

import (
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
)

func (s *service) List(reqQuery *cartModel.ListReqQuery) (resData []cartModel.ListCarts, count int64, err error) {
	cart, count, err := s.cartRepo.List(reqQuery)
	if err != nil {
		return
	}
	for i := range cart {
		cartItem, err := s.cartRepo.ListCartItem(cart[i].Id)
		if err != nil {
			return nil, 0, err
		}
		resData = append(resData, cartModel.ListCarts{
			Id:        cart[i].Id,
			UserId:    cart[i].UserId,
			CartItems: cartItem,
		})
	}

	return resData, count, nil
}
