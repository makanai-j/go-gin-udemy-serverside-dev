package main

import (
	"bytes"
	"encoding/json"
	"gin-fleamarket/dto"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"gin-fleamarket/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "Latte", Price: 300, Description: "Frothy milky coffee", SoldOut: false, UserID: 1},
		{Name: "Espresso", Price: 200, Description: "Strong black coffee", SoldOut: false, UserID: 2},
		{Name: "Cappuccino", Price: 350, Description: "Coffee with steamed milk foam", SoldOut: true, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "password1"},
		{Email: "test2@example.com", Password: "password2"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
	}

	for _, item := range items {
		if err := db.Create(&item).Error; err != nil {
			log.Fatalf("Failed to create item: %v", err)
		}
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setUpRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	//テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	//テストのセットアップ
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	CreateItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(CreateItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["item"].ID)
}

func TestCreateUnauthorized(t *testing.T) {
	//テストのセットアップ
	router := setup()

	CreateItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(CreateItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーション
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
