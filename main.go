package main

import (
	"log"

	https "github/instaShop_assessment/features/https"
	"github/instaShop_assessment/features/infastructure/database"
	"github/instaShop_assessment/features/infastructure/router"
	repository "github/instaShop_assessment/features/repository"
	usecase "github/instaShop_assessment/features/usecases"
)

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, productRepo)

	handlers := https.NewHandler(userUseCase, productUseCase, orderUseCase)
	r := router.SetupRouter(handlers)
	r.Run(":8080")
}
