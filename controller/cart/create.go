package cart

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody cartModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.cartService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyCreated, nil)
}