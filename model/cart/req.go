package cart

type (
	CartItem struct {
		ProductId uint32 `json:"product_id"`
		Quantity  uint32 `json:"quantity"`
	}

	CreateReqBody struct {
		UserId    uint32     `json:"user_id"`
		CartItems []CartItem `json:"cart_items"`
	}
)
