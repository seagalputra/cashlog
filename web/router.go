package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(SuccessResponse("Success", "pong"))
	})

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("Error starting server : %v", err)
	}
}
