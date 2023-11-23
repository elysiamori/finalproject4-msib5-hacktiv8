package repositories

import (
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{DB: db}
}

// create category
func (r *CategoryRepositoryImpl) CreateCategory(category *models.Category) (*models.Category, error) {
	err := r.DB.Create(&category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

// get categories
func (r *CategoryRepositoryImpl) GetAllCategories() ([]models.Category, error) {
	var category []models.Category
	err := r.DB.Preload("Products").Find(&category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

// get category by id
func (r *CategoryRepositoryImpl) GetCategoryById(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.Preload("Products").Where("id = ?", id).First(&category).Error

	if err != nil {
		return &category, err
	}

	return &category, nil
}

// update category
func (r *CategoryRepositoryImpl) UpdateCategory(id uint, typee string) (*models.Category, error) {
	var category models.Category
	err := r.DB.Model(&category).Where("id = ?", id).Updates(map[string]interface{}{"type": typee}).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(id string) error {
	var category models.Category
	err := r.DB.Where("id = ?", id).Delete(&category).Error

	if err != nil {
		return err
	}

	return nil
}
