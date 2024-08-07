package user

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/library/pagination"
	"github.com/AwahMarwah/Technical_Test_3/library/response"
	userModel "github.com/AwahMarwah/Technical_Test_3/model/user"
	"github.com/gin-gonic/gin"
)

func (c *controller) List(ctx *gin.Context) {
	var reqQuery userModel.ListReqQuery
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	reqQuery.Offset = pagination.Offset(&reqQuery.Limit, &reqQuery.Page)
	resData, count, err := c.userService.List(&reqQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessWithPage(ctx, http.StatusOK, "", resData, reqQuery.Page, reqQuery.Limit, count)
}
