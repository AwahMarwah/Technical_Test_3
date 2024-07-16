package cart

import (
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(reqBody *cartModel.CreateReqBody) (cartModel.Cart, error)
		List(reqQuery *cartModel.ListReqQuery) (resData []cartModel.ListCartResData, count int64, err error)
		ListCartItem(id uint32) (resData []cartModel.CartItem, err error)
		Take(selectParams []string, conditions *cartModel.Cart) (cart cartModel.Cart, err error)
		Update(id *uint32, values *map[string]any) (err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
