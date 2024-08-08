package user

import (
	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
)

func GetUserByIdHandler(c *fiber.Ctx) error {
	user_id,_:= model.GetAuthUserID(c)
	response, err := model.GetUSerByID(user_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved user",fiber.StatusOK,response)
}