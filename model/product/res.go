package porduct

type (
	ListResData struct {
		CreatedAt   string  `json:"created_at"`
		Id          uint32  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
	}
)
