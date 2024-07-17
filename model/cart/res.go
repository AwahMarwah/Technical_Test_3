package cart

type (
	CartResData struct {
		ID     uint32 `json:"id"`
		UserID uint32 `json:"user_id"`
		Status string `json:"status"`
	}

	DetailResData struct {
		Id        uint32     `json:"id"`
		UserId    uint32     `json:"user_id"`
		CartItems []CartItem `json:"cart_items"`
	}

	ListCartResData struct {
		Id     uint32 `json:"id"`
		UserId uint32 `json:"user_id"`
	}

	ListCarts struct {
		Id        uint32     `json:"id"`
		UserId    uint32     `json:"user_id"`
		CartItems []CartItem `json:"cart_items"`
	}
)
