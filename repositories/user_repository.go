package repositories

import (
	"errors"
	"html"
	"strings"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/models"
	"github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

// register user
func (r *UserRepositoryImpl) RegisterUser(user *models.User) (*models.User, error) {
	err := r.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// encrypt password before save
func (r *UserRepositoryImpl) BeforeSave(user *models.User) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Fullname = html.EscapeString(strings.TrimSpace(user.Fullname))
	return nil
}

// login user check
func (r *UserRepositoryImpl) LoginCheck(email, password string) (string, error) {
	var err error

	u := models.User{}

	err = r.DB.Model(models.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// checking role
	role := "customer"
	if u.Role == "admin" {
		role = "admin"
	}

	token, err := token.GenerateToken(u.ID, role)

	if err != nil {
		return "", err
	}

	return token, nil
}

// get user by id
func (r *UserRepositoryImpl) GetUserById(uid uint) (models.User, error) {
	var u models.User

	err := r.DB.First(&u, uid).Error

	if err != nil {
		return u, errors.New("User not found")
	}

	r.PrepareGive()

	return u, nil
}

func (r *UserRepositoryImpl) PrepareGive() {
	var user models.User
	user.Password = ""
}

// top up balance
func (r *UserRepositoryImpl) TopUpBalance(userID uint, balance int) (int, error) {
	var u models.User
	err := r.DB.First(&u, userID).Error
	if err != nil {
		return 0, err // user not found
	}

	// validate balance
	if balance < 0 {
		return u.Balance, errors.New("Topup balance must be greater than 0s")
	}

	errQ := r.DB.Model(&u).Where("id = ?", userID).Update("balance", gorm.Expr("balance + ?", balance)).Error

	if errQ != nil {
		return u.Balance, errQ
	}

	return u.Balance, nil
}
