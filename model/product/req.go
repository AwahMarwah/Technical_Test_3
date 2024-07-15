package porduct

type (
	CreateReqBody struct {
		Name        string  `binding:"required"`
		Price       float64 `binding:"required"`
		Description string
	}

	DetailReqPath struct {
		Id uint32 `binding:"required" uri:"id"`
	}

	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int `form:"page"`
	}

	UpdateReq struct {
		Body struct {
			Name        string  `binding:"required"`
			Description string  `binding:"required"`
			Price       float64 `binding:"required"`
		}
		Path struct {
			Id uint32 `binding:"required" uri:"id"`
		}
	}
)
