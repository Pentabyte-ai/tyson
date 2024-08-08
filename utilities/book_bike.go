package utilities

import (
	"net/mail"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendBail(email,id string,cost, vat,total string) error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// Compose the email
	subject := "Tyson Bikes Booking Info"
	body := "You have booked Bike ID:"+id+ "\n\n" + "Cost per hour:" +cost+"\n\n"+"vat:"+vat+"\n\n"+"total amount:"+total
	msg := []byte("Subject: " + subject + "\r\n" +
		"To: " + email + "\r\n" +
		"\r\n" +
		body)

	// Create the "from" address
	// fromAddr := mail.Address{Name: "Dancan", Address: from}
	fromAddr := mail.Address{Address: from}
	// Establish a connection to the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAddr.Address, []string{email}, msg)
	if err != nil {
		return err
	}
	return nil
}