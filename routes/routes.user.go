package routes

import (
	"Forum-Back-End/database"
	"Forum-Back-End/models"
	"Forum-Back-End/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type State struct {
	Message string
	Auth    bool
	Token   string
}

type User struct {
	ID             string `json:"id"`
	CreatedAt      time.Time
	Username       string `json:"username"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
	Email          string `json:"email"`
}

type Response struct {
	User
	State
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:             user.ID,
		Username:       user.Username,
		CreatedAt:      user.CreatedAt,
		Password:       user.Password,
		VerifyPassword: user.VerifyPassword,
		Email:          user.Email,
	}
}

func CreateResponseState(message string, auth bool, token string) State {
	return State{Message: message, Auth: auth, Token: token}
}

func CreateUser(c *fiber.Ctx) error {

	var checkField = []string{"username", "password", "verify_password", "email"}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(State{Message: "baddd", Auth: false})
	}
	if !utils.CheckField(user, checkField) {
		return c.Status(400).JSON(State{Message: "baddd", Auth: false})
	}
	if user.Password != user.VerifyPassword {
		return c.Status(400).JSON(State{Message: "Pas le meme mdp", Auth: false})
	}

	var UUID = uuid.New()
	user.ID = UUID.String()

	user.Password, _ = utils.HashPassword(user.Password)
	user.VerifyPassword, _ = utils.HashPassword(user.VerifyPassword)
	user.ID = UUID.String()

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	responseState := CreateResponseState("OK", true, "KO")

	return c.Status(200).JSON(Response{responseUser, responseState})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	var responseUsers []User
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id string, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	fmt.Println(nil, "nil")
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id", "")
	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
