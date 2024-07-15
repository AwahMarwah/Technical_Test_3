package auth

type (
	SignInReqBody struct {
		Email    string `binding:"email"`
		Password string `binding:"gte=8"`
	}
)
