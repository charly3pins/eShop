package main

import (
	"log"
	"net/http"
	"os"

	"github.com/charly3pins/eShop/application"
	"github.com/charly3pins/eShop/infrastructure/api"
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

	userRepository := postgres.UserRepository{Db: db}
	userService := application.UserService{UserRepository: userRepository}
	userHandler := api.UserHandler{UserService: userService}

	orderRepository := postgres.OrderRepository{Db: db}
	orderService := application.OrderService{OrderRepository: orderRepository}
	orderHandler := api.OrderHandler{OrderService: orderService}

	r := api.BuildRouter(userHandler, orderHandler)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
