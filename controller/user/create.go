package user

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody userModel.CreateReq
	if err := ctx.ShouldBind(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.userService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyCreated, nil)
}
