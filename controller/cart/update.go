package cart

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	cartModel "github.com/AwahMarwah/Technical_Test_3/model/cart"
	"github.com/gin-gonic/gin"
)

func (c *controller) Update(ctx *gin.Context) {
	var reqPath cartModel.UpdateReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if statusCode, err := c.cartService.Update(&reqPath); err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyUpdated, nil)
	return
}