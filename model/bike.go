package model

import (
	"errors"
	"log"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddBike(c *fiber.Ctx) error {
	bike := new(Bike)
	url, _:=utilities.GenerateUrl(c,"image")
	if err := c.BodyParser(bike); err != nil {
		return errors.New("failed to parse request data")
	}
	bike.ImageUrl=url
	err := db.Create(&bike).Error
	if err != nil{
		return errors.New("failed to add bike")
	}
	return nil
}
//update bikes
func UpdateBike(c *fiber.Ctx, bike_id uuid.UUID)(*Bike, error){
	bike := new(Bike)
	url, _:=utilities.GenerateUrl(c,"image")
	if err:=c.BodyParser(&bike); err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse request data")
	}
	bike.ImageUrl = url
	err := db.Model(&bike).Where("id = ?", bike_id).Updates(&bike).Scan(&bike).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to update bikes")
	}
	return bike,err
}
//get bikes by location
func GetBikeByLocation(location string)(*Bike,error){
	bike := new(Bike)
	if err := db.Model(&bike).Where("location = ?",location).Find(&bike).Scan((bike)).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			err_str := "no bikes found for location "+location
			return nil, errors.New(err_str)
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get bikes by location")
	}
	return bike,nil
}


