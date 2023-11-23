package repositories

import (
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"

	"gorm.io/gorm"
)

// category repo
type UserRepository interface {
	RegisterUser(user *models.User) (*models.User, error)
	BeforeSave(user *models.User) error
	LoginCheck(email, password string) (string, error)
	GetUserById(uid uint) (models.User, error)
	PrepareGive()
	TopUpBalance(uid uint, balance int) error
	CheckBalance(totalPrice int) error
	DecrementBalance(id uint, value uint, tx *gorm.DB) error
}

// category repo
type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
	UpdateCategory(uid uint, typee string) (*models.Category, error)
	DeleteCategory(id string) error
}

// product repo
type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(id uint, title string, price int, stock int, category_id int) (*models.Product, error)
	DeleteProduct(id string) error
	CheckStock(quantity int) error
	DecrementStock(id uint, quantity uint, tx *gorm.DB) error
}

// transaction repo
