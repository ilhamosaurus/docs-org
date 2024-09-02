package service

import (
	"errors"
	"os"
	"time"

	"go-templ/infra/models"
	"go-templ/infra/types"
	"go-templ/pkg/database"
	"go-templ/pkg/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type UpdateUserRequest struct {
	Email       string
	OldPassword *string
	NewPassword *string
	Name        string
}

func Login(email, password string) (*string, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	claim := &types.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := os.Getenv("SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func CreateUser(user models.RegisterRequest) error {
	db := database.DB

	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	if err := db.Create(&models.User{Email: user.Email, Password: hashPassword, Name: user.Name}).Error; err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]models.User, error) {
	db := database.DB

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db := database.DB

	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user UpdateUserRequest) (*models.User, error) {
	db := database.DB
	existingUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if user.OldPassword != nil {
		if !util.CheckPasswordHash(*user.OldPassword, existingUser.Password) {
			return nil, errors.New("old password does not match")
		}

		hashPassword, err := util.HashPassword(*user.NewPassword)
		if err != nil {
			return nil, err
		}

		if err := db.Save(&models.User{ID: existingUser.ID, Email: user.Email, Password: hashPassword, Name: user.Name}).Error; err != nil {
			return nil, err
		}

		return &models.User{ID: existingUser.ID, Email: user.Email, Password: hashPassword, Name: user.Name}, nil
	}

	if err := db.Save(&models.User{ID: existingUser.ID, Email: user.Email, Name: user.Name}).Error; err != nil {
		return nil, err
	}

	return &models.User{ID: existingUser.ID, Email: user.Email, Name: user.Name}, nil
}

func DeleteUser(id uuid.UUID) error {
	db := database.DB

	if err := db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
