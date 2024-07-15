package auth

type (
	ReqHeader struct {
		Authorization string `binding:"required"`
	}

	SignInReqBody struct {
		Email    string `binding:"email"`
		Password string `binding:"gte=8"`
	}
)
