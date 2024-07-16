package cart

type (
	CartItem struct {
		Id        uint32 `json:"id"`
		CartId    uint32 `json:"cart_id"`
		ProductId uint32 `json:"product_id"`
		Quantity  uint32 `json:"quantity"`
	}

	CreateReqBody struct {
		UserId    uint32     `json:"user_id"`
		CartItems []CartItem `json:"cart_items"`
	}

	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int    `form:"page"`
		Status string `binding:"oneof=active checkout" form:"status"`
	}

	UpdateReqPath struct {
		Id uint32 `binding:"required" uri:"id"`
	}
)
