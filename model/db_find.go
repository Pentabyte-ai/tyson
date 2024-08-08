package model

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/DANCANKARANI/tyson/database"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
var db =database.ConnectDB()
/*
finds user using phone number only
@params phone_number
*/

func UserExist(c *fiber.Ctx, email string) (bool, *User, error) {
    var existingUser User

    // Detailed logging
    log.Printf("Checking for user with pemail: %s", email)

    result := db.Where("email = ?", email).First(&existingUser)
    if result.Error != nil {
        // Log the detailed error
        log.Printf("Error finding user with email adrress %s: %v", email, result.Error)

        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", result.Error)
    }
	log.Printf("User found: %+v", existingUser)
    return true, &existingUser, nil
}

/*
provider exists
*/
func ProviderExist(c *fiber.Ctx, email string) (bool, *Provider, error) {
    var existingProvider Provider

    // Detailed logging
    log.Printf("Checking for user with email: %s", email)

    result := db.Where("email = ?", email).First(&existingProvider)
    if result.Error != nil {
        // Log the detailed error
        log.Printf("Error finding user with email adrress %s: %v", email, result.Error)

        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", result.Error)
    }
	log.Printf("User found: %+v", existingProvider)
    return true, &existingProvider, nil
}

/*
updates the reset password code in the database
@params phone_number
@params email
@params reset_code
@paarams expiration_time
*/
func AddResetCode(c *fiber.Ctx,phone_number,email,code string,exp_time time.Time) error {
	user := User{}
	db.AutoMigrate(&user)
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to get user",fiber.StatusInternalServerError)
	}

	result = db.Save(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to save data",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"code sent",fiber.StatusOK)
}
/*
finds if the user with the given email and phone number is registered
@params email
@params phone_number
*/
func FindUser(email, phoneNumber string) (User, error) {
	user := User{}
	err_str := fmt.Sprintf("user with email %s and phone number %s does not exist", email, phoneNumber)
	err := db.Where("phone_number = ? AND email = ?", phoneNumber, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			
			return user, errors.New(err_str)
		}
		return user, errors.New(err_str)
	}
	return user, nil
}


/*
finds dependants using phone number only
@params phone_number
*/
func GetAuthUserID(c *fiber.Ctx)(uuid.UUID,error){
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return uuid.Nil,errors.New("unauthorized")
	}
	user_id=*id
	return user_id.(uuid.UUID),nil
}