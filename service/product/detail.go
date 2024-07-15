package product

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

func (s *service) Detail(reqPath *productModel.DetailReqPath) (resData productModel.DetailResData, statusCode int, err error) {
	product, err := s.productRepo.Take([]string{"created_at", "description", "name", "price"}, &productModel.Product{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.ProductNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.CreatedAt = product.CreatedAt.Format("02 Jan 2006")
	resData.Description = product.Description
	resData.Name = product.Name
	resData.Price = product.Price
	return
}
