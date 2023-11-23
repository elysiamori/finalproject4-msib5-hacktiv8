package controllers

import (
	"net/http"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/dto"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/repositories"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/responses"
	token "github.com/elysiamori/finalproject4-hacktiv8-msib5/token"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionRepo repositories.TransactionRepositoryImpl
}

func NewTransactionController(transactionRepo repositories.TransactionRepositoryImpl) *TransactionController {
	return &TransactionController{TransactionRepo: transactionRepo}
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	input := dto.TransactionInput{}

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error get user id",
		})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to process request",
		})

		return
	}

	transaction, errS := tc.TransactionRepo.CreateTransaction(user_id, input)
	if errS != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errS.Error(),
		})

		return
	}

	responses := responses.CTResponse{
		Message: "You have successfully purchased the product",
		TransactionBill: responses.TransactionBill{
			TotalPrice:   transaction.TotalPrice,
			Quantity:     transaction.Quantity,
			ProductTitle: transaction.ProductTitle,
		},
	}

	c.JSON(http.StatusOK, responses)
}

func (tc *TransactionController) GetMyTransactions(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error get user id",
		})
		return
	}

	transactions, errS := tc.TransactionRepo.GetTransactionHistory(user_id)
	if errS != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error get my transaction history",
		})

		return
	}

	response := []responses.TransactionHistory{}

	for _, transaction := range transactions {
		response = append(response, responses.TransactionHistory{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: responses.Products{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      transaction.Product.Stock,
				CategoryID: transaction.Product.CategoryID,
				CreatedAt:  transaction.Product.CreatedAt,
				UpdatedAt:  transaction.Product.UpdatedAt,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TransactionController) GetAllTransactions(c *gin.Context) {
	transactions, errS := tc.TransactionRepo.GetAllTransactionHistory()
	if errS != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error get all transaction history",
		})

		return
	}

	response := []responses.TransactionHistories{}

	for _, transaction := range transactions {
		response = append(response, responses.TransactionHistories{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: responses.Products{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      transaction.Product.Stock,
				CategoryID: transaction.Product.CategoryID,
				CreatedAt:  transaction.Product.CreatedAt,
				UpdatedAt:  transaction.Product.UpdatedAt,
			},
			User: responses.Users{
				ID:        transaction.User.ID,
				Email:     transaction.User.Email,
				Fullname:  transaction.User.Fullname,
				Balance:   transaction.User.Balance,
				CreatedAt: transaction.User.CreatedAt,
				UpdatedAt: transaction.User.UpdatedAt,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}
