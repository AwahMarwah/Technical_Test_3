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
	ReqPath struct {
		Id uint32 `binding:"required" uri:"id"`
	}
	SeedReq struct {
		Email    string `validate:"email"`
		Password string `validate:"gte=8"`
	}
)
