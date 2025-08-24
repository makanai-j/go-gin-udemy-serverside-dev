package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// postgreSQLの疎通確認
	// >>>> ここから
	user := "postgres"
	pass := "postgres"
	db := "postgres"
	host := "localhost"
	port := "5432"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Errorf("db ping failed: %w", err))
	}
	fmt.Println("DB ping OK")
	_, err = pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS items(id serial primary key, name text not null)`)
	if err != nil {
		panic(err)
	}
	_, err = pool.Exec(ctx, `INSERT INTO items(name) VALUES($1)`, "hello")
	if err != nil {
		panic(err)
	}
	var cnt int
	if err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM items`).Scan(&cnt); err != nil {
		panic(err)
	}
	fmt.Println("items count =", cnt)
	// <<<<<< ここまで

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
	router.PUT("/items/:id", itemController.Update)
	router.DELETE("/items/:id", itemController.Delete)
	router.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
