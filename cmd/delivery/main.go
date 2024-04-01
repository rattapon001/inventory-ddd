package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	delivery_router "github.com/rattapon001/inventory-ddd/api/routers/delivery"
	"github.com/rattapon001/inventory-ddd/internal/delivery/app"
	eventhandler "github.com/rattapon001/inventory-ddd/internal/delivery/infra/event_handler"
	delivery_memory "github.com/rattapon001/inventory-ddd/internal/delivery/infra/memory"
)

func main() {

	err := godotenv.Load("./configs/.env.delivery")
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	repository := delivery_memory.NewDeliveryMemoryRepository()
	publisherRepo := eventhandler.NewMemoryOutboxRepository()
	publisher := eventhandler.NewPublisher(publisherRepo)

	deliveryUseCase := app.NewDeliveryUseCase(repository, publisher)
	delivery_router.InitRouter(router, &deliveryUseCase)

	router.Run(":" + port)

}
