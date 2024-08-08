package provider

import (
	"log"

	"github.com/DANCANKARANI/tyson/database"
	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
var db = database.ConnectDB()
func CreateProviderAccount(c *fiber.Ctx) error {
	//generating new id
	id := uuid.New()
	provider := model.Provider{}
	if err := c.BodyParser(&provider); err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c, "failed to create account", fiber.StatusInternalServerError)
	}

	//validate email address
	_, err := utilities.ValidateEmail(provider.Email)
	if err != nil {
		return utilities.ShowError(c, err.Error(), fiber.StatusInternalServerError)
	}
	//Check if user exist
	userExist, _, err := model.ProviderExist(c, provider.Email)
	if err != nil {
		return utilities.ShowError(c, err.Error(), fiber.StatusInternalServerError)
	}
	if userExist {
		errStr := "user with this email " + provider.Email + " already exists"
		return utilities.ShowError(c, errStr, fiber.StatusConflict)
	}
	
	//comapare passwords
	if provider.ConfirmPassword != provider.Password {
		return utilities.ShowError(c, "passwords do not match", fiber.StatusForbidden)
	}

	//hash password
	hashed_password, _ := utilities.HashPassword(provider.Password)

	providerModel := model.Provider{ID: id, FullName: provider.FullName, Email: provider.Email, Password: hashed_password}
	//create user
	err = db.Create(&providerModel).Error
	if err != nil {
		log.Fatal(err.Error())
		return utilities.ShowError(c, "failed to add data to the database", fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c, "account created successfully", fiber.StatusOK)
}