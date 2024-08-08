package model

import (
	"errors"
	"log"
	"strconv"

	//"strconv"

	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddBike(c *fiber.Ctx, providerID uuid.UUID) error {
    bike := new(Bike)

    // Parse the form data
    form, err := c.MultipartForm()
    if err != nil {
        log.Printf("Error parsing multipart form: %v", err)
        return errors.New("failed to parse form data")
    }

    // Get the file from the form
   

    // Save the file
    url, err := utilities.SaveFile(c, "image")
    if err != nil {
        log.Printf("Error saving file: %v", err)
        return errors.New("failed to save file")
    }

    // Get other form fields
    bike.Location = form.Value["location"][0] 
// Convert price string to float64
	priceStr := form.Value["cost"][0]
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Println("Error converting price:", err)
	
	}

// Assign the parsed value to bike.Price
bike.Price = price

	bike.CalculateVAT(16)
	
    bike.ID = uuid.New()
    bike.ProviderID = providerID
    bike.ImageUrl = url

    if err := db.Create(&bike).Error; err != nil {
        log.Printf("Error adding bike to database: %v", err)
        return errors.New("failed to add bike")
    }

    return nil
}

//update bikes
func UpdateBike(c *fiber.Ctx, bike_id uuid.UUID)(*Bike, error){
	bike := new(Bike)
	url, _:=utilities.SaveFile(c,"image")
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
func GetBikeByLocation(location string) (*[]Bike, error) {
    bikes := new([]Bike)

    // Use SOUNDEX to match similar-sounding names
    query := "SELECT * FROM bikes WHERE SOUNDEX(location) = SOUNDEX(?)"

    if err := db.Raw(query, location).Scan(&bikes).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println(err.Error())
            errStr := "no bikes found for location " + location
            return nil, errors.New(errStr)
        }
        log.Println(err.Error())
        return nil, errors.New("failed to get bikes by location")
    }

    return bikes, nil
}




func (p *Bike)CalculateVAT(vatRate float64){
	p.Vat = p.Price*vatRate/100
	p.Total = p.Price+p.Vat
}