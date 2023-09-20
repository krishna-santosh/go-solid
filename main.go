package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"go-solid/ui"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(ui.UI),
		PathPrefix: "dist",
	}))

	app.Listen(":9000")
}
