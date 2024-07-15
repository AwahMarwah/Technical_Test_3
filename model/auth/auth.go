package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId uint32
	*jwt.StandardClaims
}
