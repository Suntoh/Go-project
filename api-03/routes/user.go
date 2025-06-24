package routes

import (
	"api-03/database"
	"api-03/models"

	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	//this is not the midel User, see this as serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusCreated).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	if err := database.Database.Db.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	responseUsers := []UserResponse{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	// if err := database.Database.Db.First(user, id).Error; err != nil {
	// 	return err
	// }
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id,err := c.ParamsInt("id")
	user := models.User{}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user ID")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}