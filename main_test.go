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
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, UserID: 1},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, UserID: 2},
	}
	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})
	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	// Request, Path, Body
	req, _ := http.NewRequest("GET", "/items", nil)
	// テストの実行
	router.ServeHTTP(w, req)
	// レスポンスの検証
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)
	// assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	// テストのセットアップ
	router := setup()

	// 認証
	token, err := services.CreateToken(1, "test1@example.com")
	// エラー確認
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Create Test",
	}
	// json encode
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	// Request, Path, Body
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// テストの実行
	router.ServeHTTP(w, req)
	// レスポンスの検証
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)
	// assertion
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

// func TestFindById(t *testing.T) {
// 	// テストのセットアップ
// 	router := setup()

// 	w := httptest.NewRecorder()
// 	// Request, Path, Body
// 	req, _ := http.NewRequest("GET", "/items/1", nil)
// 	// テストの実行
// 	router.ServeHTTP(w, req)
// 	// レスポンスの検証
// 	var res map[string][]models.Item
// 	json.Unmarshal(w.Body.Bytes(), &res)
// 	// assertion
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
// 	assert.Equal(t, 0, len(res["data"]))
// }
