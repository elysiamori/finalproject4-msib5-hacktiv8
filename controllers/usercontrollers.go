package controllers

import (
	"net/http"
	"strconv"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/dto"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/repositories"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/responses"
	token "github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo repositories.UserRepositoryImpl
}

func (uc *UserController) CurrentUser(c *gin.Context) {

	// mengambil token id
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := uc.UserRepo.GetUserById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    u,
	})
}

func (uc *UserController) UserRegister(c *gin.Context) {

	// input data dari user
	input := dto.RegisterInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	u := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
		Role:     "customer",
	}

	// enkripsi password di database
	if err := uc.UserRepo.BeforeSave(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// save data user ke database
	_, errS := uc.UserRepo.RegisterUser(&u)
	if errS != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errS.Error(),
		})
		return
	}

	response := responses.UserResponse{
		ID:        u.ID,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Password:  input.Password,
		Balance:   u.Balance,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusCreated, response)

}

func (uc *UserController) LoginUser(c *gin.Context) {

	//input data login dari user
	input := dto.LoginInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// validasi email dan password
	token, err := uc.UserRepo.LoginCheck(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password is incorrect",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (uc *UserController) TopUpBalance(c *gin.Context) {

	if err := token.TokenValid(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	topupData := dto.TopUpInput{}

	err = c.ShouldBindJSON(&topupData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// validasi jumlah top-up
	if topupData.Balance > 100000000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Maximum top-up is Rp 100.000.000",
		})
		return
	}

	// balance updated to database
	balanceUp, errID := uc.UserRepo.TopUpBalance(userID, topupData.Balance)
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to top up",
		})
		return
	}

	updatedBalance := topupData.Balance + balanceUp

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Your balance has been succesfully updated to Rp " + strconv.Itoa(updatedBalance),
	})

}
