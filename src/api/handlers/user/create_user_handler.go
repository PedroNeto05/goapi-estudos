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
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	password, err := hashPassword(userRequest.Password)

	if err != nil {
		logger.Errorf("Erro ao fazer o hash da senha: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Email:    userRequest.Email,
		Password: password,
	}

	err = userusecase.CreateUserUseCase(&user)

	if err != nil {
		logger.Errorf("Erro na criação do usuario: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
