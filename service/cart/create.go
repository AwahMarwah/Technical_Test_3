package cart

import (
	"errors"

	"github.com/AwahMarwah/Technical_Test_3/common"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"gorm.io/gorm"
)

func (s *service) Create(reqBody *cartModel.CreateReqBody) (err error) {
	_, err = s.userRepo.Take([]string{"id"}, &userModel.User{Id: reqBody.UserId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(common.UserNotFound)
		}
		return err
	}
	_, err = s.cartRepo.Create(reqBody)
	if err != nil {
		return err
	}
	return
}