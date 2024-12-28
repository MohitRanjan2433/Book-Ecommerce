package main

import (
	"bookecom/config"
	"bookecom/database"
	"bookecom/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	database.ConnectDB(&config)
	database.RunMigrations(database.DB)

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", 
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))
	

	apiGroup := app.Group("/v1")

	routes.BookRoutes(apiGroup)
	routes.AuthRoutes(apiGroup)	

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}

	fmt.Printf("PORT: ", port)

	log.Fatal(app.Listen(":" + config.Port))
}