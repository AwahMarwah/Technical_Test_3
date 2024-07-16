package cart

import cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"

func (r *repo) List(reqQuery *cartModel.ListReqQuery) (resData []cartModel.ListCartResData, count int64, err error) {
	resData = make([]cartModel.ListCartResData, 0)
	return resData, count, r.db.Select("id, user_id").Model(&cartModel.Cart{}).Where("status = ?", "active").Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
