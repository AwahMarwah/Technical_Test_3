package product

import (
	"gorm.io/gorm"
	productRepo "github.com/AwahMarwah/Technical_Test_3/repository/product"
	"github.com/AwahMarwah/Technical_Test_3/service/product"
)

type controller struct {
	productService product.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{productService: product.NewService(productRepo.NewRepo(db))}
}
