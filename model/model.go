package model

import "github.com/google/uuid"

type User struct {
	ID 				uuid.UUID 	`json:"id" gorm:"type:varchar(36)"`
	FullName		string 		`json:"full_name" gorm:"type:varchar(255)"`
	Email			string		`json:"email" gorm:"type:varchar(255)"`
	Password		string		`json:"password" gorm:"type:varchar(255)"`
	ConfirmPassword	string		`json:"confirm_password" gorm:"type:varchar(255)"`
}
type Provider struct{
	ID 				uuid.UUID 	`json:"id" gorm:"type:varchar(36)"`
	BusnessName		string		`json:"business_name" gorm:"type:varchar(255)"`
	FullName		string 		`json:"full_name" gorm:"type:varchar(255)"`
	Email			string		`json:"email" gorm:"type:varchar(255)"`
	Password		string		`json:"password" gorm:"type:varchar(255)"`
	ConfirmPassword	string		`json:"confirm_password" gorm:"type:varchar(255)"`
	BusinessDescription	string	`json:"business_description" gorm:"type:varchar(255)"`
	Bike            []Bike     	`gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}

type Bike struct{
	ID 			uuid.UUID 	`json:"id" gorm:"type:varchar(36)"`
	ProviderID  uuid.UUID	`json:"provider_id" gorm:"type:varchar(36); default:NULL"`
	ImageUrl	string		`json:"image_url" gorm:"type:varchar(1024)"`
	Location	string 		`json:"location" gorm:"type:varchar(255)"`
	Price		float64		`json:"price" gorm:"type:decimal(13,4)"`
	Vat			float64		`json:"vat" gorm:"type:decimal(13,4)"`
	Total		float64		`json:"total" gorm:"type:decimal(13,4)"`
	Provider 	Provider	`json:"provider"`
}

type Plan struct{
	ID 			uuid.UUID 	`json:"id" gorm:"type:varchar(36)"`
	Type		string		`json:"type" gorm:"type:varchar(255)"`
	TotalBikes	int			`json:"total_bikes" gorm:"type:int"`
	Instrument	string		`json:"instrument" gorm:"type:varchar(255)"`
	Place		string		`json:"place" gorm:"type:varchar(255)"`
}