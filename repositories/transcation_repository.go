package repositories

import (
	"errors"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/dto"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/responses"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{DB: db}
}

func (r *TransactionRepositoryImpl) CreateTransaction(userID uint, request dto.TransactionInput) (*responses.TransactionBill, error) {

	product := models.Product{}
	tB := responses.TransactionBill{}
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&product, request.ProductID).Error; err != nil {
			return errors.New("product not found")
		}

		if request.Quantity > product.Stock {
			return errors.New("insufficient stock")
		}

		user := models.User{}
		if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
			return errors.New("user not found")
		}

		totalPrice := product.Price * request.Quantity
		if user.Balance < totalPrice {
			return errors.New("insufficient balance")
		}

		if err := tx.Model(&product).Update("stock", gorm.Expr("stock - ?", request.Quantity)).Error; err != nil {
			return errors.New("failed to update product stock")
		}

		if err := tx.Model(&user).Update("balance", gorm.Expr("balance - ?", totalPrice)).Error; err != nil {
			return errors.New("failed to update user balance")
		}

		category := models.Category{}
		if err := tx.Where("id = ?", product.CategoryID).First(&category).Error; err != nil {
			return errors.New("category not found")
		}

		if err := tx.Model(&category).Update("sold_product", gorm.Expr("sold_product + ?", request.Quantity)).Error; err != nil {
			return errors.New("failed to update category sold")
		}

		tH := responses.TransactionHistory{
			Quantity:   request.Quantity,
			TotalPrice: totalPrice,
			UserID:     userID,
			ProductID:  request.ProductID,
		}

		err := tx.Create(&tH).Error
		if err != nil {
			return errors.New("failed to create transaction history")
		}

		tB.TotalPrice = totalPrice
		tB.Quantity = request.Quantity
		tB.ProductTitle = product.Title

		return nil

	})

	return &tB, err
}

func (r *TransactionRepositoryImpl) GetTransactionHistory(userID uint) ([]responses.TransactionHistory, error) {
	var tH []responses.TransactionHistory

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Product").Where("user_id = ?", userID).Find(&tH).Error; err != nil {
			return errors.New("transaction history not found")
		}

		if len(tH) == 0 {
			return errors.New("transaction history not found")
		}

		return nil
	})

	return tH, err
}

func (r *TransactionRepositoryImpl) GetAllTransactionHistory() ([]responses.TransactionHistories, error) {
	t := []responses.TransactionHistories{}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Product").Preload("User").Find(&t).Error; err != nil {
			return errors.New("transaction history not found")
		}

		if len(t) == 0 {
			return errors.New("transaction history not found")
		}

		return nil
	})

	return t, err
}
