package main

import (
	"fmt"
	handler "pdfprocessor/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the web directory
	router.Static("/", "./web")

	router.POST("/svg-to-pdf", handler.ConvertSVGToPDF)

	port := "8080"
	fmt.Println("[INFO] Server running on port", port)
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("[ERROR] Failed to start server:", err)
	}
}
