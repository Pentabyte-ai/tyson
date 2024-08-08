package user

import (
	"fmt"
	"time"

	"github.com/DANCANKARANI/tyson/middleware"
	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
)
type ResponseUser struct{
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}
type loginResponse struct {
	Token string `json:"token"`
}

func Login(c *fiber.Ctx)error{
	user := model.User{}
	if err := c.BodyParser(&user); err !=nil {
		return utilities.ShowError(c,"failed to login",fiber.StatusInternalServerError)
	}

	//check of user exist
	userExist,existingUser,_:= model.UserExist(c,user.Email)
	if ! userExist {
		return utilities.ShowError(c,"user does not exist",fiber.StatusNotFound)
	}
	
	//compare password
	err :=utilities.CompareHashAndPassowrd(existingUser.Password,user.Password)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusForbidden)
			 
	}
	exp :=time.Hour*24
	//generating token
	tokenString,err := middleware.GenerateToken(middleware.Claims{UserID: &existingUser.ID,Role: "user"},exp)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//set token cookie 
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // Same duration as the token
		HTTPOnly: true, // Important for security, prevents JavaScript access
		Secure:   true, // Use secure cookies in production
		Path:     "/",  // Make the cookie available on all routes
	})
	response_user:=loginResponse{
		Token: tokenString,
	}
	return utilities.ShowSuccess(c,"successfully logged in",fiber.StatusOK,response_user)	
}

func Logout(c *fiber.Ctx) error {
	tokenString,err :=utilities.GetJWTToken(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusUnauthorized)
	}
	fmt.Println(tokenString)
	err = middleware.InvalidateToken(tokenString)
	if err != nil {
		return utilities.ShowError(c,"failed to invalidate the token",fiber.StatusInternalServerError)
	}
	
	return utilities.ShowMessage(c,"successfully logged out",fiber.StatusOK)
}