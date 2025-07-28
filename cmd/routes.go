package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicobellanich/go_postgreSQL_docker/handlers"
)

func setupRoutes(app *fiber.App) {

	// list all
	app.Get("/", handlers.ListFacts)

	// create one
	app.Post("/fact", handlers.CreateFact)
}
