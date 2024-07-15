package health

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) Check(ctx *gin.Context) {
	if err := c.healthService.Check(); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyChecked, nil)
}