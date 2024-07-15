package user

import userModel "github.com/AwahMarwah/Technical_Test_3/model/user"

func (r *repo) Update(id *uint32, values *map[string]any) (err error) {
	return r.db.Model(userModel.User{Id: *id}).Updates(values).Error
}
