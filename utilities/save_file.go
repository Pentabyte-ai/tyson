package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)

type ImgurResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

// SaveFile uploads the file to Imgur and returns its URL.
func SaveFile(c *fiber.Ctx, fieldName string) (string, error) {
	// Parse the form data to get the file
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Read the file content
	fileContent, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	// Upload the file to Imgur
	imageURL, err := uploadToImgur(fileContent, filepath.Ext(file.Filename))
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

// uploadToImgur uploads the image bytes to Imgur and returns the image URL.
func uploadToImgur(fileContent []byte, fileExt string) (string, error) {
	clientId := os.Getenv("IMGUR_CLIENT_ID")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "upload"+fileExt)
	if err != nil {
		return "", err
	}
	part.Write(fileContent)
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", clientId))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var imgurResp ImgurResponse
	if err := json.Unmarshal(respBody, &imgurResp); err != nil {
		return "", err
	}

	if !imgurResp.Success {
		return "", fmt.Errorf("failed to upload image: status %d", imgurResp.Status)
	}

	return imgurResp.Data.Link, nil
}

// Middleware to serve static files from the 'uploads' directory (if needed)
func StaticFileMiddleware(app *fiber.App) {
	app.Static("/uploads", "./uploads")
}
