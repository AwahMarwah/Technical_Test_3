package router

import (
	"github.com/AwahMarwah/Technical_Test_3/controller/auth"
	"github.com/AwahMarwah/Technical_Test_3/controller/cart"
	"github.com/AwahMarwah/Technical_Test_3/controller/health"
	"github.com/AwahMarwah/Technical_Test_3/controller/product"
	"github.com/AwahMarwah/Technical_Test_3/controller/root"
	"github.com/AwahMarwah/Technical_Test_3/controller/user"
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
		idGroup := productGroup.Group(":id")
		{
			idGroup.GET("", authorize(db.GormDb), productController.Detail)
			idGroup.PUT("", authorize(db.GormDb), productController.Update)
		}
	}
	cartGroup := router.Group("cart")
	{
		cartController := cart.NewController(db.GormDb)
		cartGroup.POST("", authorize(db.GormDb), cartController.Create)
		cartGroup.GET("", authorize(db.GormDb), cartController.List)
		idGroup := cartGroup.Group(":id")
		{
			idGroup.GET("", authorize(db.GormDb), cartController.Detail)
			idGroup.PUT("", authorize(db.GormDb), cartController.Update)
		}
	}
	userGroup := router.Group("user")
	{
		userController := user.NewController(db.GormDb)
		userGroup.POST("", authorize(db.GormDb), userController.Create)
		userGroup.GET("", authorize(db.GormDb), userController.List)
		idGroup := userGroup.Group(":id")
		{
			idGroup.GET("", authorize(db.GormDb), userController.Detail)
		}
	}

	return router.Run()
}
