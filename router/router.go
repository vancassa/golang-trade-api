package router

import (
	"trade-api/controllers"
	"trade-api/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	adminRouter := router.Group("/auth")
	{
		adminRouter.POST("/register", controllers.AdminRegister)
		adminRouter.POST("/login", controllers.AdminLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetAllProducts)
		// productRouter.GET("/:productUUID", controllers.GetProductById)

		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productUUID", middleware.ProductAuthorization(), controllers.UpdateProduct)
		// productRouter.DELETE("/:productUUID", middleware.ProductAuthorization(), controllers.DeleteProduct)
	}

	variantRouter := router.Group("/products/variants")
	{
		variantRouter.GET("/", controllers.GetAllVariants)
		// productRouter.GET("/:variantUUID", controllers.GetVariantById)

		variantRouter.Use(middleware.Authentication())
		variantRouter.POST("/", controllers.CreateVariant)
		variantRouter.PUT("/:variantUUID", middleware.VariantAuthorization(), controllers.UpdateVariant)
		// variantRouter.DELETE("/:variantUUID", middleware.VariantAuthorization(), controllers.DeleteVariant)
	}

	return router
}
