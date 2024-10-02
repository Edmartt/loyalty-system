package main

import (
	"log"
	"os"

	_ "github.com/Edmartt/loyalty-system/docs"
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

	commerceRepository := adapters.NewCommerceRepository(pgConnection)

	commerceService := services.NewCommerceService(commerceRepository)
	commerceHandlers := gin.CommerceHandlers{
		CommerceService: *commerceService,
	}

	branchRepository := adapters.NewBranchRepository(pgConnection)
	branchService := services.NewBranchService(branchRepository)
	branchHandler := gin.BranchHandler{
		BranchService: *branchService,
	}

	httpServer := gin.HTTPServer{
		CustomerHandler: customerHandler,
		CommerceHandler: commerceHandlers,
		BranchHandler:   branchHandler,
	}

	if err := httpServer.RunServer(os.Getenv("HTTP_PORT")); err != nil {
		panic(err)
	}
}
