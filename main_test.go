package main

import (
	"gin-fleamarket/models"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
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
}
