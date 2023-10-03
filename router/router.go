package router

import (
	"go-solid/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/notes", handlers.CreateNote)
	app.Get("/notes/:id", handlers.GetNote)
	app.Get("/notes", handlers.GetAllNotes)
	app.Patch("/notes/:id", handlers.UpdateNote)
	app.Delete("/notes/:id", handlers.DeleteNote)
}
