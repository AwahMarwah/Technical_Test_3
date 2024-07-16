package cart

type (
	CartResData struct {
		ID     uint32 `json:"id"`
		UserID uint32 `json:"user_id"`
		Status string `json:"status"`
	}
)
