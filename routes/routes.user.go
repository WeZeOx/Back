package routes

import (
	"Forum-Back-End/database"
	"Forum-Back-End/structures"
	"Forum-Back-End/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	structures.User
	structures.State
}

func CreateResponseUser(user structures.User) structures.User {
	return structures.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Password:  user.Password,
		Email:     user.Email,
	}
}

func CreateResponseUserWithPost(user structures.User, PostArr []structures.Post) structures.User {
	return structures.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Password:  user.Password,
		Email:     user.Email,
		Post:      PostArr,
	}
}

func CreateResponseState(message string, auth bool, token string) structures.State {
	return structures.State{Message: message, Auth: auth, Token: token}
}

func CreateUser(c *fiber.Ctx) error {

	var checkFieldUserArray = []string{"username", "password", "verify_password", "email"}
	var user structures.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Missing fields", Auth: false})
	}
	if !utils.CheckFieldUser(user, checkFieldUserArray) {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Missing fields", Auth: false})
	}
	if user.Password != user.VerifyPassword {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Not the same password / verify password", Auth: false})
	}
	if utils.EmailExist(user) {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "email already exist", Auth: false})
	}
	if utils.UsernameExist(user) {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "username already exist", Auth: false})
	}

	user.ID = uuid.New().String()

	user.Password, _ = utils.HashPassword(user.Password)
	user.VerifyPassword, _ = utils.HashPassword(user.VerifyPassword)

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	token := utils.CreateToken(user)
	responseState := CreateResponseState("OK", true, token)

	return c.Status(fiber.StatusOK).JSON(Response{responseUser, responseState})
}

func GetUsers(c *fiber.Ctx) error {
	var post []structures.Post

	var users []structures.User
	var responseUsersAndPost []structures.User

	database.Database.Db.Find(&users)

	for _, user := range users {
		database.Database.Db.Where("user_id = ?", user.ID).Find(&post)
		responseUsersAndPost = append(responseUsersAndPost, CreateResponseUserWithPost(user, post))
	}

	return c.Status(fiber.StatusOK).JSON(responseUsersAndPost)
}

func findUser(id string, user *structures.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id", "")
	var user structures.User
	var post []structures.Post

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Where("user_id = ?", id).Find(&post)
	response := CreateResponseUserWithPost(user, post)

	return c.Status(fiber.StatusOK).JSON(response)
}

func CheckFieldLogin(user Login, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)
	if err != nil {
		return false
	}

	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}

func LoginUser(c *fiber.Ctx) error {
	var login Login
	var user structures.User
	var checkFieldLoginArray = []string{"email", "password"}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Missing fields", Auth: false})
	}
	if !CheckFieldLogin(login, checkFieldLoginArray) {
		return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Missing fields", Auth: false})
	}

	database.Database.Db.Where("email = ?", login.Email).Find(&user)

	if utils.CheckPasswordHash(login.Password, user.Password) {
		token := utils.CreateToken(user)

		return c.JSON(Response{
			User: structures.User{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				Username:  user.Username,
				Email:     user.Email,
			},
			State: structures.State{
				Message: "Authentified",
				Auth:    true,
				Token:   token,
			}})
	}

	return c.Status(fiber.StatusBadRequest).JSON(structures.State{Message: "Email / Password Incorrect", Auth: false})
}
