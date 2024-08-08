package utilities
import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	password_bytes := []byte(password)
	hashed_password_bytes, err := bcrypt.GenerateFromPassword(password_bytes,bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashed_password_bytes), nil
}