package porduct

type (
	CreateReqBody struct {
		Name        string  `binding:"required"`
		Price       float64 `binding:"required"`
		Description string
	}
)
