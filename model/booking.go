package model

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type BookingDetail struct{
	Email		string
	BikeID		uuid.UUID
	Cost		float64
	Vat			float64
	Total		float64
}

func BookBike(user_id, bike_id uuid.UUID) (*BookingDetail,error) {
	user := new(User)
	err:=db.First(&user,"id=?",user_id).Scan(&user).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get user")
	}
	bike := new(Bike)
	err = db.First(&bike,"id = ?",bike_id).Scan(&bike).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to get bike details")
	}
	response := BookingDetail{
		Email: user.Email,
		Cost: bike.Price,
		Vat: bike.Vat,
		Total: bike.Total,
	}
	return &response,nil
}