package cart

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"gorm.io/gorm"
)

func (s *service) Update(reqPath *cartModel.UpdateReqPath) (statusCode int, err error) {
	if _, err = s.cartRepo.Take([]string{"id"}, &cartModel.Cart{Id: reqPath.Id}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New(common.CartNotFound)
		}
		return http.StatusInternalServerError, err
	}
	values := map[string]any{
		"status": "checkout",
	}
	if err = s.cartRepo.Update(&reqPath.Id, &values); err != nil {
		return http.StatusInternalServerError, err
	}
	return
}