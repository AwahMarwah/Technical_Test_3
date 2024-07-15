package health

import healthRepo "github.com/AwahMarwah/Technical_Test_3/repository/health"

type (
	IService interface {
		Check() (err error)
	}

	service struct {
		healthRepo healthRepo.IRepo
	}
)

func NewService(healthRepo healthRepo.IRepo) IService {
	return &service{healthRepo: healthRepo}
}