package cart

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

func (s *service) List(reqQuery *cartModel.ListReqQuery) (resData []cartModel.ListCarts, count int64, err error) {
	resData = make([]cartModel.ListCarts, 0)
	cart, count, err := s.cartRepo.List(reqQuery)
	if err != nil {
		return nil, 0, err
	}
	for i := range cart {
		cartItem, err := s.cartRepo.ListCartItem(cart[i].Id)
		if err != nil {
			return nil, 0, err
		}
		for j := range cartItem {
			product, err := s.productRepo.Find([]string{"price", "description"}, &productModel.Product{Id: cartItem[j].ProductId})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return resData, http.StatusNotFound, errors.New(common.ProductNotFound)
				}
				return resData, http.StatusInternalServerError, err
			}
			if len(product) > 0 {
				cartItem[j].Price = product[0].Price
			}
		}
		// Inisialisasi CartItems sebagai array kosong jika cartItem kosong
		if cartItem == nil {
			cartItem = []cartModel.CartItem{}
		}
		resData = append(resData, cartModel.ListCarts{
			Id:        cart[i].Id,
			UserId:    cart[i].UserId,
			CartItems: cartItem,
		})
	}

	return resData, count, nil
}
