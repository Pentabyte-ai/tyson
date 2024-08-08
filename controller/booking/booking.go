package booking

import (
	"strconv"
	"github.com/DANCANKARANI/tyson/model"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func BookingHandler(c *fiber.Ctx)error{
	user_id, _:= model.GetAuthUserID(c)
	bike_id,_:=uuid.Parse(c.Params("id"))
	response, err := model.BookBike(user_id,bike_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(), fiber.StatusInternalServerError)
	}
	cost := strconv.FormatFloat(response.Cost, 'f', 1, 64) // "123.5"
    vat := strconv.FormatFloat(response.Vat, 'f', 2, 64) // "123.46"
    total := strconv.FormatFloat(response.Total, 'f', 3, 64) // "123.457"
	err=utilities.SendBail(response.Email,bike_id.String(),cost,vat,total)
	if err != nil{
		return utilities.ShowError(c,"failed to send bailing infomation", fiber.StatusInsufficientStorage)
	}
	return utilities.ShowMessage(c,"bailing information sent successfully",fiber.StatusOK)
}