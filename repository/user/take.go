package user

import userModel "github.com/AwahMarwah/Technical_Test_3/model/user"

func (r *repo) Take(selectParams []string, conditions *userModel.User) (user userModel.User, err error) {
	return user, r.db.Select(selectParams).Take(&user, conditions).Error
}
