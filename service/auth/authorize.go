package auth

import (
	"errors"
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/encrypt"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"gorm.io/gorm"
)

func (s *service) Authorize(token *string) (userId uint32, statusCode int, err error) {
	tokenRaw, claims, err := encrypt.Parse(*token)
	if err != nil {
		return userId, http.StatusUnauthorized, err
	}
	userId = uint32(claims["UserId"].(float64))
	user, err := s.userRepo.Take([]string{"token"}, &userModel.User{Id: userId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userId, http.StatusUnauthorized, errors.New(common.UserNotFound)
		}
		return userId, http.StatusInternalServerError, err
	}
	if tokenRaw != user.Token {
		return userId, http.StatusUnauthorized, errors.New(common.UserHasSignedOut)
	}
	return userId, http.StatusOK, nil
}
