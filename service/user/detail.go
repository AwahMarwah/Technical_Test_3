package user

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"gorm.io/gorm"
)

func (s *service) Detail(reqPath *userModel.ReqPath) (resData userModel.DetailUserResData, statusCode int, err error) {
	user, err := s.userRepo.Take([]string{"id", "created_at", "email"}, &userModel.User{Id: reqPath.Id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusNotFound, errors.New(common.UserNotFound)
		}
		return resData, http.StatusInternalServerError, err
	}
	resData.ID = user.Id
	resData.CreatedAt = user.CreatedAt.Format("02 January 2006")
	resData.Email = user.Email
	return
}
