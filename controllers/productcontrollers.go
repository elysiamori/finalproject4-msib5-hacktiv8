package controllers

import (
	"net/http"
	"strconv"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/dto"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/helpers"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/repositories"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/responses"
	token "github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductRepo repositories.ProductRepositoryImpl
}

func (pc *ProductController) CreateProduct(c *gin.Context) {

	input := dto.ProductInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Price > 50000000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Maximum price is 50.000.000",
		})
		return
	}

	product := models.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}

	_, errS := pc.ProductRepo.CreateProduct(&product)
	if errS != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errS.Error(),
		})
		return
	}

	response := responses.ProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, response)
}

// get all products
func (pc *ProductController) GetProducts(c *gin.Context) {

	products, err := pc.ProductRepo.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productRes []responses.ProductResponse

	for _, product := range products {
		productRes = append(productRes, responses.ProductResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, productRes)
}

// update product
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	idConv, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingProduct, err := pc.ProductRepo.GetProductById(uint(idConv))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := dto.ProductInput{}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productUp, err := pc.ProductRepo.UpdateProduct(uint(idConv), input.Title, input.Price, input.Stock, int(input.CategoryID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := &responses.UpdateProductID{
		Product: responses.ProductResponseID{
			ID:         existingProduct.ID,
			Title:      productUp.Title,
			Price:      helpers.FormatMoney(productUp.Price),
			Stock:      productUp.Stock,
			CategoryID: productUp.CategoryID,
			CreatedAt:  existingProduct.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  productUp.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	c.JSON(http.StatusOK, response)

}

// delete products
func (pc *ProductController) DeleteProduct(c *gin.Context) {
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

	err = pc.ProductRepo.DeleteProduct(strconv.Itoa(int(idConv)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted",
	})
}
