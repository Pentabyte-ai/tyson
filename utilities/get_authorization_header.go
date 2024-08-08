package utilities

import (
	"errors"
	"strings"
	"github.com/gofiber/fiber/v2"
)

/*
gets the jwt token from authorization header
*/

func GetJWTToken(c *fiber.Ctx)(string,error){
	authHeader := c.Get(fiber.HeaderAuthorization)
	if authHeader == "" {
		
		return "", errors.New("failed to get the token from the header")
	}
	//split the token into "Bearer" and token
	tokenString :=strings.TrimSpace(strings.TrimPrefix(authHeader,"Bearer"))

	if tokenString == ""{
		
		return "", errors.New("failed to split the token")
	}
	return tokenString, nil
}