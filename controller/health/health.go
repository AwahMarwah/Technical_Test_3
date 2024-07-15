package health

import (
	"database/sql"

	"github.com/AwahMarwah/Technical_Test_3/service/health"
	healthRepo "github.com/AwahMarwah/Technical_Test_3/repository/health"
)

type controller struct {
	healthService health.IService
}

func NewController(db *sql.DB) *controller {
	return &controller{healthService: health.NewService(healthRepo.NewRepo(db))}
}
