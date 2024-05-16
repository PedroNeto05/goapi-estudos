package userusecase

import (
	"errors"
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUserUseCase(login *models.LoginRequest) (*string, error) {
	err := isAValidLogin(login)

	if err != nil {
		return nil, err
	}

	userExist, err := userrepository.GetUserByEmail(login.Email)

	if !(errors.Is(err, gorm.ErrRecordNotFound)) && err != nil {
		return nil, err
	}

	if userExist == nil {
		return nil, errors.New("Invalid email or password. Please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(login.Password))

	if err != nil {
		return nil, errors.New("Invalid email or password. Please try again")
	}

	token, err := createToken(userExist)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func isAValidLogin(login *models.LoginRequest) error {

	if login.Email == "" && login.Password == "" {
		return errors.New("the body is empty or poorly formatted")
	}

	if login.Email == "" {
		return errors.New("the email is required")
	}

	if login.Password == "" {
		return errors.New("the password is required")
	}

	return nil

}

func createToken(user *models.User) (*string, error) {

	secretKey := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return nil, err
	}

	return &token, nil
}
