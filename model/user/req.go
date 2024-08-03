package user

type (
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
