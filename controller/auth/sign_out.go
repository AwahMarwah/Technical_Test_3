package auth

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) SignOut(ctx *gin.Context) {
	if err := c.authService.SignOut(ctx.MustGet("userId").(uint32)); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullySignedOut, nil)
}
