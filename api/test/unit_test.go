package test

import (
	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

type AppConfig struct {
	Email string `yaml:"email"`
	Name string `yaml:"name"`
}

func TestReadYaml(t *testing.T) {
	b, err := os.ReadFile("config.yml")
	if err != nil {
		t.Fatal(err)
	}
	var config AppConfig
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(config.Email)
	log.Println(config.Name)
}

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
