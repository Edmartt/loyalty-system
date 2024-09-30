package main

import (
	"log"
	"os"

	"github.com/Edmartt/loyalty-system/internal/adapters"
	"github.com/Edmartt/loyalty-system/internal/adapters/database"
	"github.com/Edmartt/loyalty-system/internal/adapters/http/gin"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading env vars: %v", err)
	}

	pgConnection, err := database.NewPostgresConn()

	if err != nil {
		panic(err.Error())
	}

	customerRepository := adapters.NewCustomersRepository(pgConnection)

	customerService := services.NewCustomerService(customerRepository)

	customerHandler := gin.CustomerHandlers{
		CustomerService: *customerService,
	}

	httpServer := gin.HTTPServer{
		Handler: customerHandler,
	}

	if err := httpServer.RunServer(os.Getenv("HTTP_PORT")); err != nil {
		panic(err)
	}
}
