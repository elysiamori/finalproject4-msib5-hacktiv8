package routers

import (
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/controllers"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(usercontroller *controllers.UserController, categorycontroller *controllers.CategoryController,
	productcontroller *controllers.ProductController, transactioncontroller *controllers.TransactionController) *gin.Engine {

	r := gin.Default()

	// ------- user register and login ---------
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", usercontroller.UserRegister)
		userRouter.POST("/login", usercontroller.LoginUser)
		userRouter.PATCH("/topup", middlewares.JwtAuthMiddleWares(), usercontroller.TopUpBalance)
	}

	// ------- middlewares get user --------
	protected := r.Group("/profile")
	{
		protected.GET("/users", middlewares.JwtAuthMiddleWares(), usercontroller.CurrentUser)
	}

	// ------ admin middlewares categories ------
	protectedAdmin := r.Group("/")
	{
		protectedAdmin.Use(middlewares.JwtAuthMiddleWares())
		protectedAdmin.GET("/categories", middlewares.OnlyAdminAuth(), categorycontroller.GetAllCategories)
		protectedAdmin.GET("/categories/:id", middlewares.OnlyAdminAuth(), categorycontroller.GetCategoryById)
		protectedAdmin.POST("/categories", middlewares.OnlyAdminAuth(), categorycontroller.CreateCategory)
		protectedAdmin.PATCH("/categories/:id", middlewares.OnlyAdminAuth(), categorycontroller.UpdateCategory)
		protectedAdmin.DELETE("/categories/:id", middlewares.OnlyAdminAuth(), categorycontroller.DeleteCategory)
	}

	// ------ admin middlewares products ------
	protectedAdminProduct := r.Group("/")
	{
		protectedAdminProduct.Use(middlewares.JwtAuthMiddleWares())
		protectedAdminProduct.GET("/products", productcontroller.GetProducts)
		protectedAdminProduct.POST("/products", middlewares.OnlyAdminAuth(), productcontroller.CreateProduct)
		protectedAdminProduct.PUT("/products/:id", middlewares.OnlyAdminAuth(), productcontroller.UpdateProduct)
		protectedAdminProduct.DELETE("/products/:id", middlewares.OnlyAdminAuth(), productcontroller.DeleteProduct)
	}

	// ------ admin middlewares transactions ------
	protectedAdminTransaction := r.Group("/transactions")
	{
		protectedAdminTransaction.Use(middlewares.JwtAuthMiddleWares())
		protectedAdminTransaction.POST("/", transactioncontroller.CreateTransaction)
		protectedAdminTransaction.GET("/my-transactions", transactioncontroller.GetMyTransactions)
		protectedAdminTransaction.GET("/user-transactions", middlewares.OnlyAdminAuth(), transactioncontroller.GetAllTransactions)
	}

	return r
}
