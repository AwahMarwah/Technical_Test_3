package product

import (
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(product *productModel.Product) (err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}