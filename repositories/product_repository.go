package repositories

import (
	"errors"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{DB: db}
}

// create product
func (r *ProductRepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	err := r.DB.Create(&product).Error

	if err != nil {
		return nil, err
	}

	if product.Stock < 5 {
		return nil, errors.New("Minimum stock is 5")
	}

	if product.Price < 0 {
		return nil, errors.New("Minimum price is 0")
	}

	return product, nil
}

// get all products
func (r *ProductRepositoryImpl) GetAllProducts() ([]models.Product, error) {
	var product []models.Product
	err := r.DB.Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

// get product by id
func (r *ProductRepositoryImpl) GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.Where("id = ?", id).First(&product).Error

	if err != nil {
		return &product, err
	}

	return &product, nil
}

// update product
func (r *ProductRepositoryImpl) UpdateProduct(id uint, title string, price int, stock int, category_id int) (*models.Product, error) {
	var product models.Product

	err := r.DB.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "price": price, "stock": stock, "category_id": category_id}).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// delete product
func (r *ProductRepositoryImpl) DeleteProduct(id string) error {
	var product models.Product
	err := r.DB.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
}
