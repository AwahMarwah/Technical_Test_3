package cart

import (
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(reqBody *cartModel.CreateReqBody) (cartModel.Cart, error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
