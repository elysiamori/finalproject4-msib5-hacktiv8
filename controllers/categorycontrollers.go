package controllers

import (
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/dto"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/repositories"

	"net/http"
	"strconv"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/responses"
	token "github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryRepo repositories.CategoryRepositoryImpl
}

// add new category
func (cc *CategoryController) CreateCategory(c *gin.Context) {

	input := dto.CategoryInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category := models.Category{
		Type: input.Type,
	}

	categoryAdd, errS := cc.CategoryRepo.CreateCategory(&category)
	if errS != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errS.Error(),
		})
		return
	}

	response := responses.CategoryResponseAdd{
		ID:          categoryAdd.ID,
		Type:        categoryAdd.Type,
		SoldProduct: categoryAdd.SoldProduct,
		CreatedAt:   categoryAdd.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, response)

}

// get all categories
func (cc *CategoryController) GetAllCategories(c *gin.Context) {

	categories, err := cc.CategoryRepo.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var categoryResponse []responses.CategoryResponse
	for _, category := range categories {
		var products []responses.GetProductResponse
		for _, product := range category.Products {
			products = append(products, responses.GetProductResponse{
				ID:        product.ID,
				Title:     product.Title,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		categoryResponse = append(categoryResponse, responses.CategoryResponse{

			ID:          category.ID,
			Type:        category.Type,
			SoldProduct: category.SoldProduct,
			CreatedAt:   category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   category.UpdatedAt.Format("2006-01-02 15:04:05"),
			Products:    products,
		})
	}

	c.JSON(http.StatusOK, categoryResponse)

}

// get category by id
func (cc *CategoryController) GetCategoryById(c *gin.Context) {

	id := c.Params.ByName("id")
	idConv, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := cc.CategoryRepo.GetCategoryById(uint(idConv))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var products []responses.GetProductResponse
	for _, product := range category.Products {
		products = append(products, responses.GetProductResponse{
			ID:        product.ID,
			Title:     product.Title,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	responseCategory := responses.CategoryResponse{
		ID:          category.ID,
		Type:        category.Type,
		SoldProduct: category.SoldProduct,
		CreatedAt:   category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   category.UpdatedAt.Format("2006-01-02 15:04:05"),
		Products:    products,
	}

	c.JSON(http.StatusOK, responseCategory)
}

// update category
func (cc *CategoryController) UpdateCategory(c *gin.Context) {

	id := c.Params.ByName("id")
	idConv, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := token.TokenValid(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	existingCategory, err := cc.CategoryRepo.GetCategoryById(uint(idConv))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	input := dto.CategoryInput{}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	categoryUp, err := cc.CategoryRepo.UpdateCategory(uint(idConv), input.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	categoryResponse := responses.CategoryResponseUpdate{
		ID:          existingCategory.ID,
		Type:        categoryUp.Type,
		SoldProduct: existingCategory.SoldProduct,
		UpdatedAt:   categoryUp.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, categoryResponse)
}

// delete category
func (cc *CategoryController) DeleteCategory(c *gin.Context) {

	id := c.Params.ByName("id")
	idConv, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := token.TokenValid(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	categoryId, err := cc.CategoryRepo.GetCategoryById(uint(idConv))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	err = cc.CategoryRepo.DeleteCategory(strconv.Itoa(int(categoryId.ID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been sucessfully deleted",
	})
}
