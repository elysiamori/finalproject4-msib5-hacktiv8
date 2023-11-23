package main

import (
	"os"
	
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/controllers"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/database"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/repositories"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/routers"
)

func main() {
	db, err := database.Config()
	if err != nil {
		panic(err)
	}

	// user
	userRepository := repositories.NewUserRepository(db)
	userController := controllers.UserController{UserRepo: *userRepository}

	// category
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryController := controllers.CategoryController{CategoryRepo: *categoryRepository}

	// product
	productRepository := repositories.NewProductRepository(db)
	productController := controllers.ProductController{ProductRepo: *productRepository}

	// transaction
	transactionRepository := repositories.NewTransactionRepository(db)
	transactionController := controllers.TransactionController{TransactionRepo: *transactionRepository}

	r := routers.StartApp(&userController, &categoryController, &productController, &transactionController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
