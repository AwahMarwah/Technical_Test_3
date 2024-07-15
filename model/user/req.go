package user

type SeedReq struct {
	Email    string `validate:"email"`
	Password string `validate:"gte=8"`
}
