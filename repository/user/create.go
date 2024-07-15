package user

import userModel "github.com/AwahMarwah/Technical_Test_3/model/user"

func (r *repo) Create(user *userModel.User) (err error) {
	return r.db.Create(user).Error
}
