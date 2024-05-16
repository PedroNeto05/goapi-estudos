package userhandler

import (
	userusecase "goApi/api/usecase/user"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var userRequest *models.User

	err := c.BodyParser(&userRequest)

	if err != nil {
		logger.Errorf("Error parsing information: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	password, err := hashPassword(userRequest.Password)

	if err != nil {
		logger.Errorf("Error hashing the password: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := models.User{
		ID:        uuid.New(),
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  password,
	}

	err = userusecase.CreateUserUseCase(&user)

	if err != nil {
		logger.Errorf("Error creating the user: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
