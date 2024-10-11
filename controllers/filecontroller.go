package controllers

import (
	"go-bucket/usecases"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	usecase *usecases.FileUseCase
}

func CreateFileController(usecase *usecases.FileUseCase) FileController {
	return FileController{usecase: usecase}
}

func (filecontroller *FileController) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get file", "error": err.Error()})
		return
	}
	defer file.Close()

	fileData, err := filecontroller.usecase.UploadFileUseCase(file, header.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to save file", "error": err.Error()})
		return
	}
	
	c.JSON(201, gin.H{
		"status": "success",
		"fileId": fileData.Id,
		"createdAtUTC": fileData.CreatedAtUTC,
		"key": fileData.Key,
	})
}

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := "data/" + filename

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(filePath)
}