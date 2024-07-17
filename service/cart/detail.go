package cart

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

func (s *service) Detail(reqPath *cartModel.ReqPath) (resData cartModel.DetailResData, statusCode int, err error) {
	cart, err := s.cartRepo.Take([]string{"id", "user_id"}, &cartModel.Cart{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.CartNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.CartItems = make([]cartModel.CartItem, 0)
	cartItem, err := s.cartRepo.ListCartItem(cart.Id)
	if err != nil {
		return resData, http.StatusInternalServerError, err
	}
	for i := range cartItem {
		product, err := s.productRepo.Find([]string{"price", "description"}, &productModel.Product{Id: cartItem[i].ProductId})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return resData, http.StatusNotFound, errors.New(common.ProductNotFound)
			}
			return resData, http.StatusInternalServerError, err
		}
		if len(product) > 0 {
			cartItem[i].Price = product[0].Price
		}
	}
	resData.Id = cart.Id
	resData.UserId = cart.UserId
	resData.CartItems = cartItem
	return
}
