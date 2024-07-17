package product

import (
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(product *productModel.Product) (err error)
		List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListResData, count int64, err error)
		Take(selectParams []string, conditions *productModel.Product) (product productModel.Product, err error)
		Update(id *uint32, values *map[string]any) (err error)
		Find(selectParams []string, conditions *productModel.Product) (resData []productModel.Product, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
