package user

type (
	DetailUserResData struct {
		ID        uint32 `json:"id"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}
	ListUser struct {
		ID        uint32 `json:"id"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}
)
