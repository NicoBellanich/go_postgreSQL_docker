package main

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	// list all
	app.Get("/", handlers.ListFacts)

	// create one
	app.Post("/fact", handlers.CreateFact)
}
