package user

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/library/response"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"github.com/gin-gonic/gin"
)

func (c *controller) Detail(ctx *gin.Context) {
	var reqPath userModel.ReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.userService.Detail(&reqPath)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, "", resData)
	return
}
