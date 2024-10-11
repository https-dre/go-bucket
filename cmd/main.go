package main

import (
	"go-bucket/controllers"
	"go-bucket/repositories"
	"go-bucket/storage"
	"go-bucket/usecases"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != os.Getenv("BACKEND_KEY") && token != os.Getenv("BACKEND_MOBILE") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Next()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	router := gin.Default()
	router.Use(AuthMiddleware)

	db_sqlite := storage.Init()

	fileRepo := repositories.CreateFileRepository(db_sqlite)
	fileUsecase := usecases.CreateFileUseCase(&fileRepo)
	fileController := controllers.CreateFileController(&fileUsecase)

	router.POST("/upload", fileController.UploadFile)

	router.Run(":8080")
}


