package product

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/common"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"github.com/gin-gonic/gin"
)

func (c *controller) Update(ctx *gin.Context) {
	var req productModel.UpdateReq
	if err := ctx.ShouldBindUri(&req.Path); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if statusCode, err := c.productService.Update(&req); err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyUpdated, nil)
	return
}