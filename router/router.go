package router

import (
	"github.com/AwahMarwah/Technical_Test_3/controller/auth"
	"github.com/AwahMarwah/Technical_Test_3/controller/health"
	"github.com/AwahMarwah/Technical_Test_3/controller/product"
	"github.com/AwahMarwah/Technical_Test_3/controller/root"
	"github.com/AwahMarwah/Technical_Test_3/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	rootController := root.NewController()
	router.GET("", rootController.Index)
	healthController := health.NewController(db.SqlDb)
	router.GET("health", healthController.Check)
	authGroup := router.Group("auth")
	{
		authController := auth.NewController(db.GormDb)
		authGroup.POST("sign-in", authController.SignIn)
		authGroup.DELETE("sign-out", authorize(db.GormDb), authController.SignOut)
	}
	productGroup := router.Group("product")
	{
		productController := product.NewController(db.GormDb)
		productGroup.GET("", authorize(db.GormDb), productController.List)
		productGroup.POST("", authorize(db.GormDb), productController.Create)
	}

	return router.Run()
}