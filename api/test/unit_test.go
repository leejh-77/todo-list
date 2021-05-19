package test

import (
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