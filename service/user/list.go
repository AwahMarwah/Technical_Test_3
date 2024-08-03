package user

import (
	"time"

	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
)

func (s *service) List(reqQuery *userModel.ListReqQuery) (resData []userModel.ListUser, count int64, err error) {
	resData = make([]userModel.ListUser, 0)
	user, count, err := s.userRepo.List(reqQuery)
	if err != nil {
		return nil, 0, err
	}
	for i := range user {
		createdAt, err := time.Parse(time.RFC3339, user[i].CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		createdAtFormatted := createdAt.Format("02 January 2006")
		resData = append(resData, userModel.ListUser{
			ID: user[i].ID,
			Email: user[i].Email,
			CreatedAt: createdAtFormatted,
		})
	}
	return
}