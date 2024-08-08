package utilities

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"path/filepath"
)

// SaveFile handles saving the uploaded file and generating its URL.
func SaveFile(c *fiber.Ctx, fieldName string) (string, error) {
	// Parse the form data to get the file
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	// Generate a unique file name
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	savePath := filepath.Join("uploads", newFilename)

	// Save the file to the server
	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	// Generate URL for the file
	fileURL := fmt.Sprintf("/uploads/%s", newFilename)
	return fileURL, nil
}

// Middleware to serve static files from the 'uploads' directory
func StaticFileMiddleware(app *fiber.App) {
	app.Static("/uploads", "./uploads")
}
