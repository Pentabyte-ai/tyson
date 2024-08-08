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
type ResBike struct{
	ID 			uuid.UUID 	`json:"id" gorm:"type:varchar(36)"`
	ProviderID  uuid.UUID	`json:"provider_id" gorm:"type:varchar(36); default:NULL"`
	ImageUrl	string		`json:"image_url" gorm:"type:varchar(1024)"`
	Location	string 		`json:"location" gorm:"type:varchar(255)"`
	Price		float64		`json:"price" gorm:"type:decimal(13,4)"`
	Vat			float64		`json:"vat" gorm:"type:decimal(13,4)"`
	Total		float64		`json:"total" gorm:"type:decimal(13,4)"`
	Owner      string    `json:"owner"`
}
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
func GetBikeByLocation( location string) (*[]ResBike, error) {
	bikes := new([]ResBike)

	// Use SOUNDEX to match similar-sounding names and JOIN to get provider details
	query := `
        SELECT bikes.id, bikes.provider_id, bikes.image_url, bikes.location, bikes.price, bikes.vat, bikes.total,
               providers.full_name AS owner
        FROM bikes
        LEFT JOIN providers ON bikes.provider_id = providers.id
        WHERE SOUNDEX(bikes.location) = SOUNDEX(?)
    `

	if err := db.Raw(query, location).Scan(&bikes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			errStr := "no bikes found for location " + location
			return nil, errors.New(errStr)
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get bikes by location")
	}

	// Log result to verify the owner field
	for _, bike := range *bikes {
		log.Printf("Bike ID: %s, Owner: %s\n", bike.ID, bike.Owner)
	}

	return bikes, nil
}

func GetAllBikes() (*[]ResBike, error) {
    bikes := new([]ResBike)

    // Query to get all bikes with provider details
    query := `
        SELECT bikes.id, bikes.provider_id, bikes.image_url, bikes.location, bikes.price, bikes.vat, bikes.total,
               providers.full_name AS owner
        FROM bikes
        LEFT JOIN providers ON bikes.provider_id = providers.id
    `

    if err := db.Raw(query).Scan(&bikes).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("No bikes found")
            return nil, errors.New("no bikes found")
        }
        log.Println(err.Error())
        return nil, errors.New("failed to get all bikes")
    }

    // Log result to verify the owner field
    for _, bike := range *bikes {
        log.Printf("Bike ID: %s, Owner: %s\n", bike.ID, bike.Owner)
    }

    return bikes, nil
}


func (p *Bike)CalculateVAT(vatRate float64){
	p.Vat = p.Price*vatRate/100
	p.Total = p.Price+p.Vat
}