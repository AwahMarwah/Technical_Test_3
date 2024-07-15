package router

import (
	"net/http"

	authModel "github.com/AwahMarwah/Technical_Test_3/model/auth"
	userRepo "github.com/AwahMarwah/Technical_Test_3/repository/user"

	"github.com/AwahMarwah/Technical_Test_3/library/response"
	"github.com/AwahMarwah/Technical_Test_3/service/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func authorize(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqHeader authModel.ReqHeader
		if err := ctx.ShouldBindHeader(&reqHeader); err != nil {
			response.Error(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		authService := auth.NewService(userRepo.NewRepo(db))
		userId, statusCode, err := authService.Authorize(&reqHeader.Authorization)
		if err != nil {
			response.Error(ctx, statusCode, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
	}
}
