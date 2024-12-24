package router

import (
	"github.com/gin-gonic/gin"

	https "github/instaShop_assessment/features/https"
	middleware "github/instaShop_assessment/features/middleware"
)

func SetupRouter(h *https.Handler) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Product routes
		products := auth.Group("/products")
		products.Use(middleware.AdminOnly())
		{
			products.POST("/", h.CreateProduct)
			products.GET("/", h.GetProducts)
			products.PUT("/:id", h.UpdateProduct)
			products.DELETE("/:id", h.DeleteProduct)
		}

		// Order routes
		orders := auth.Group("/orders")
		{
			orders.POST("/", h.CreateOrder)
			orders.GET("/", h.GetUserOrders)
			orders.PUT("/:id/cancel", h.CancelOrder)
			orders.PUT("/:id/status", middleware.AdminOnly(), h.UpdateOrderStatus)
		}
	}

	return r
}
