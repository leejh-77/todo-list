package test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	err := os.Setenv("jwt-secret", "jwt-secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(os.Getenv("jwt-secret"))
}

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.Table("products").Create(&Product{Code: "D42", Price: 100})
	db.Table("products").Create(&Product{Code: "D42", Price: 100})

	var products []Product
	db.Table("products").Find(&products)

	t.Log(products)
}