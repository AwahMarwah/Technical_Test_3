package user

import userModel "github.com/AwahMarwah/Technical_Test_3/model/user"

func (r *repo) List(reqQuery *userModel.ListReqQuery) (resData []userModel.ListUser, count int64, err error) {
	resData = make([]userModel.ListUser, 0)
	return resData, count, r.db.Select("id, email, created_at").Model(&userModel.User{}).Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
