package porduct

type (
	DetailResData struct {
		CreatedAt   string  `json:"created_at"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	ListResData struct {
		CreatedAt   string  `json:"created_at"`
		Id          uint32  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
)
