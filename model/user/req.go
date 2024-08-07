package user

type (
	CreateReq struct {
		Email             string `form:"email" binding:"required,email"`
		EncryptedPassword string `form:"encrypted_password"`
	}
	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int `form:"page"`
	}
	SeedReq struct {
		Email    string `validate:"email"`
		Password string `validate:"gte=8"`
	}
)
