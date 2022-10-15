package common

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return "", err
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "uploaded_file/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return "", err
	}

	return newFileName, nil
}

func UploadFiles(c *gin.Context) ([]string, error) {
	form, err := c.MultipartForm()
	imageAddresses := []string{}

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return imageAddresses, err
	}

	files := form.File["files"]

	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension

		if err := c.SaveUploadedFile(file, "uploaded_file/"+newFileName); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return imageAddresses, err
		}

		imageAddresses = append(imageAddresses, "uploaded_file/"+newFileName)
	}

	return imageAddresses, nil
}
