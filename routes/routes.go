package routes

import (
	"backend_started/controllers"
	"backend_started/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/products", controllers.GetAllProducts)
		protected.POST("/products", controllers.CreateProduct)
		protected.GET("/users", controllers.GetAllUsers)
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/payments", controllers.GetAllPayments)
		protected.POST("/payments", controllers.CreatePayment)
		protected.GET("/transactions", controllers.GetAllTransactions)
		protected.POST("/transactions", controllers.CreateTransaction)
	}
}
