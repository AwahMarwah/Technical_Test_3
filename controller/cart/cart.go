package cart

import (
	"github.com/AwahMarwah/Technical_Test_3/service/cart"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"
	productRepo "github.com/AwahMarwah/Technical_Test_3/repository/product"
	cartRepo "github.com/AwahMarwah/Technical_Test_3/repository/cart"
	"gorm.io/gorm"
)

type controller struct {
	cartService cart.IService

}

func NewController(db *gorm.DB) *controller {
	return &controller{cartService: cart.NewService(cartRepo.NewRepo(db), userRepo.NewRepo(db), productRepo.NewRepo(db))}
}