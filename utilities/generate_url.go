package utilities

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GenerateUrl reads the uploaded file, stores it in the 'uploads' folder, and returns a URL.
func GenerateUrl(c *fiber.Ctx,file string) (string, error) {
	// Get the file from the request
	fileHeader, err := c.FormFile(file)
	if err != nil {
		log.Println("Error retrieving file from form:", err.Error(), " ",file)
		return "", fiber.NewError(fiber.StatusBadRequest, "failed to retrieve file")
	}

	// Generate a unique file name using UUID and preserve the original file extension
	fileExt := filepath.Ext(fileHeader.Filename)
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)

	// Define the upload folder
	uploadFolder := "./uploads"
	if _, err := os.Stat(uploadFolder); os.IsNotExist(err) {
		err = os.Mkdir(uploadFolder, os.ModePerm)
		if err != nil {
			log.Println("Error creating upload directory:", err)
			return "", fiber.NewError(fiber.StatusInternalServerError, "failed to create upload directory")
		}
	}

	// Create the file path
	filePath := filepath.Join(uploadFolder, newFileName)

	// Save the file to the specified path
	if err := c.SaveFile(fileHeader, filePath); err != nil {
		log.Println("Error saving file:", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, "failed to save file")
	}

	// Generate a URL for the file
	// Assuming the file server serves files from the uploads directory
	baseURL := "http://localhost:3000" // Change this to your server's base URL
	fileURL := fmt.Sprintf("%s/uploads/%s", baseURL, newFileName)

	return fileURL, nil
}