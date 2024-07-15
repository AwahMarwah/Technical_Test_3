package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/encrypt"
	authModel "github.com/AwahMarwah/Technical_Test_3/model/auth"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

func (s *service) SignIn(reqBody *authModel.SignInReqBody) (resData authModel.SignInResData, statusCode int, err error) {
	user, err := s.userRepo.Take([]string{"encrypted_password", "id"}, &userModel.User{Email: reqBody.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusBadRequest, errors.New(common.EmailOrPasswordIsIncorrect)
		}
		return resData, http.StatusInternalServerError, err
	}
	if err = encrypt.CompareHashAndPassword(&user.EncryptedPassword, &reqBody.Password); err != nil {
		return resData, http.StatusBadRequest, errors.New(common.EmailOrPasswordIsIncorrect)
	}
	claims := authModel.Claims{
		UserId: user.Id,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
	}
	resData.Token, err = encrypt.NewTokenWithClaims(claims)
	if err != nil {
		return resData, http.StatusInternalServerError, err
	}
	if err = s.userRepo.Update(&user.Id, &map[string]any{"token": resData.Token}); err != nil {
		return resData, http.StatusInternalServerError, err
	}
	return resData, http.StatusOK, nil
}