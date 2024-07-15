package root

import (
	"net/http"

	"github.com/AwahMarwah/Technical_Test_3/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) Index(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, "", nil)
}
