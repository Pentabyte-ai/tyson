package bike

import (
	"log"

	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddBikeHandler(c *fiber.Ctx)error{
	provider_id,err :=model.GetAuthUserID(c)
	if err != nil{
		log.Println(err.Error())
	}
	log.Println(provider_id)
	err = model.AddBike(c,provider_id)
	if err != nil{
		return utilities.ShowError(c, err.Error(), fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully added bikes", fiber.StatusOK)
}

func UpdateBikeHandler(c *fiber.Ctx)error{
	bike_id,err:= uuid.Parse(c.Params("id"))
	if err != nil{
		log.Println(err.Error())
		return err
	}
	log.Println(bike_id)
	response, err := model.UpdateBike(c,bike_id)
	if err != nil{
		return utilities.ShowError(c, err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated bikes",fiber.StatusOK,response)
}

func GetBikeByLocationHandler(c *fiber.Ctx)error{
	location := c.Query("location")
	if location ==""{
		log.Println("empty location")
		return utilities.ShowError(c,"location is empty",fiber.StatusInternalServerError)
	}
	response, err := model.GetBikeByLocation(location)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c, "successfully retrieved bikes by location",fiber.StatusOK,response)
}

func GetAllBikesHandler(c *fiber.Ctx)error{
	response, err := model.GetAllBikes()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfull retrieved bikes",fiber.StatusOK,response)
}