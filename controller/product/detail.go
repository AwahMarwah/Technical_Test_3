package product

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/library/response"
	productModel "github.com/AwahMarwah/Technical_Test_3/model/product"
	"github.com/gin-gonic/gin"
)

func (c *controller) Detail(ctx *gin.Context) {
	var reqPath productModel.DetailReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.productService.Detail(&reqPath)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, "", resData)
}
