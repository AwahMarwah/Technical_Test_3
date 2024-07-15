package auth

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	authModel "github.com/AwahMarwah/Technical_Test_3/model/auth"
	"github.com/gin-gonic/gin"
)

func (c *controller) SignIn(ctx *gin.Context) {
	var reqBody authModel.SignInReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.authService.SignIn(&reqBody)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, common.SuccessfullySignedIn, resData)
}