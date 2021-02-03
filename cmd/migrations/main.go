package main

import (
	"os"

	"github.com/charly3pins/eShop/domain"
	"github.com/charly3pins/eShop/infrastructure/postgres"
)

func main() {
	connectionOptions := postgres.ConnectionOptions{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DBNAME"),
	}

	db, err := postgres.NewConnection(connectionOptions)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// TODO create models for DB and add the FK there
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.Product{})
}
