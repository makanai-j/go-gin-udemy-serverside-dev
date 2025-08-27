package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setUpRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	router.Use(cors.Default())
	itemRouter := router.Group("/items")
	itemRouterWithAuth := router.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := router.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.LogIn)

	return router
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	router := setUpRouter(db)

	// items := []models.Item{
	// 	{ID: 1, Name: "Latte", Price: 300, Description: "Frothy milky coffee", SoldOut: false},
	// 	{ID: 2, Name: "Espresso", Price: 200, Description: "Strong black coffee", SoldOut: false},
	// 	{ID: 3, Name: "Cappuccino", Price: 350, Description: "Coffee with steamed milk foam", SoldOut: true},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)

	router.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
