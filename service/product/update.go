package product

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

func (s *service) Update(req *productModel.UpdateReq) (statusCode int, err error) {
	if _, err = s.productRepo.Take([]string{"id"}, &productModel.Product{Id: req.Path.Id}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New(common.ProductNotFound)
		}
		return http.StatusInternalServerError, err
	}
	values := map[string]any{
		"name": req.Body.Name,
		"description": req.Body.Description,
		"price": req.Body.Price,
	}
	if err = s.productRepo.Update(&req.Path.Id, &values); err != nil {
		return http.StatusInternalServerError, err
	}
	return
}