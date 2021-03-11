package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/seagalputra/cashlog/pkg/db"
	"github.com/seagalputra/cashlog/pkg/transaction"
	"github.com/seagalputra/cashlog/pkg/user"
)

// Server definition
type Server struct {
	TransactionHandler transaction.Handler
	UserHandler        user.Handler
}

// Start server
func (s *Server) Start() {
	app := fiber.New()

	config, err := GetConfig(os.Getenv("APP_ENV"), ".")
	if err != nil {
		panic(err)
	}

	conn, err := db.Connect(config.DBUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Pong")
	})

	err = app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		fmt.Printf("Error starting server : %v", err)
	}
}
