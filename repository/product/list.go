package product

import productModel "github.com/AwahMarwah/Technical_Test_3/model/product"

func (r *repo) List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListResData, count int64, err error) {
	resData = make([]productModel.ListResData, 0)
	return resData, count, r.db.Select("TO_CHAR(created_at, 'DD Mon YYYY') AS created_at, id, description, price, name").Model(productModel.Product{}).Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}