package middleware

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/DANCANKARANI/tyson/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID *uuid.UUID `json:"user_id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
//the function loads .env and return the secretKey
func LoadSecretKey()string{
	err := godotenv.Load(".env")
	if err != nil {
		return err.Error()
	}
	my_secret_key := os.Getenv("MY_SECRET_KEY")
	return my_secret_key
}

/*
Generates a JWT token
@params claims *Claims 
@params expiration time
*/
func GenerateToken(claims Claims,expiration_time time.Duration) (string, error) {
	my_secret_key := LoadSecretKey()
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiration_time).Unix(),
		Issuer:    "qvp",
		
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(my_secret_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*
Validates the token string
@params tokenString
*/ 
func ValidateToken(tokenString string)(*Claims,error){
	token,err := jwt.ParseWithClaims(tokenString, &Claims{},func(token *jwt.Token) (interface{}, error) {
		return []byte("MY_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	claims,ok := token.Claims.(*Claims); 
	if ! ok{
		return nil, errors.New("invalid user token")
	}
	isRevoked, err := database.RedisClient().SIsMember(context.Background(),"revoked_tokens",tokenString).Result()
	if err != nil{
		return nil, err
	}
	if isRevoked{
		return nil, errors.New("user token is revoked")
	}
	return claims,nil
}

/*
Invalidates token when logged out
@params tokenString
*/
func InvalidateToken(tokenString string)error{
	err := database.RedisClient().SAdd(context.Background(), "revoked_token", tokenString).Err()
	if err != nil{
		return err
	}
	return nil
}

/*
gets the users id from the token
@params claims *Claims
*/
func GetAuthUserID(c *fiber.Ctx,claims *Claims)(*uuid.UUID,error){
	if claims == nil{
		return nil, errors.New("unauthorized user denied. user details not found")
	}
	//extract the user ID from the claims
	userClaimID := claims.UserID
	if userClaimID ==nil{
		return nil, errors.New("unauthorized user denied. user details not found")
	}
	return userClaimID, nil
}