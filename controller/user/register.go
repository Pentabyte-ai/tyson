package user

import (
	"log"

	"github.com/DANCANKARANI/tyson/database"
	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
var db = database.ConnectDB()
func CreateUserAccount(c *fiber.Ctx) error {
	db.AutoMigrate(&model.User{})
	//generating new id
	id := uuid.New()
	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c, "failed to create account", fiber.StatusInternalServerError)
	}

	//validate email address
	_, err := utilities.ValidateEmail(user.Email)
	if err != nil {
		return utilities.ShowError(c, err.Error(), fiber.StatusInternalServerError)
	}
	//Check if user exist
	userExist, _, err := model.UserExist(c, user.Email)
	if err != nil {
		return utilities.ShowError(c, err.Error(), fiber.StatusInternalServerError)
	}
	if userExist {
		errStr := "user with this email " + user.Email + " already exists"
		return utilities.ShowError(c, errStr, fiber.StatusConflict)
	}
	
	//comapare passwords
	if user.ConfirmPassword != user.Password {
		return utilities.ShowError(c, "passwords do not match", fiber.StatusForbidden)
	}

	//hash password
	hashed_password, _ := utilities.HashPassword(user.Password)

	userModel := model.User{ID: id, FullName: user.FullName, Email: user.Email, Password: hashed_password}
	//create user
	err = db.Create(&userModel).Error
	if err != nil {
		log.Fatal(err.Error())
		return utilities.ShowError(c, "failed to add data to the database", fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c, "account created successfully", fiber.StatusOK)
}