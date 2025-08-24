package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "Latte", Price: 300, Description: "Frothy milky coffee", SoldOut: false},
		{ID: 2, Name: "Espresso", Price: 200, Description: "Strong black coffee", SoldOut: false},
		{ID: 3, Name: "Cappuccino", Price: 350, Description: "Coffee with steamed milk foam", SoldOut: true},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	router := gin.Default()
	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindById)
	router.POST("/items", itemController.Create)
	router.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
