package model

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

func GetUSerByID(user_id uuid.UUID) (*User, error) {
	user := new(User)
	if err := db.First(&user,"id = ?",user_id).Scan(&user).Error; err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to get user")
	}
	return user, nil
}