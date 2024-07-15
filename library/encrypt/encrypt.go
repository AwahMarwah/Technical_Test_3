package encrypt

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hashedPassword, password *string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
}

func GenerateFromPassword(password *string) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.MinCost)
	if err != nil {
		return
	}
	*password = string(encryptedPassword)
	return
}

func NewTokenWithClaims(claims jwt.Claims) (token string, err error) {
	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimsToken.SignedString([]byte("simple_ecommerce"))
}