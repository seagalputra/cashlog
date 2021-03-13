package server

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/seagalputra/cashlog/pkg/config"
	"github.com/seagalputra/cashlog/pkg/transaction"
	"github.com/seagalputra/cashlog/pkg/user"
)

// Server definition
type Server struct {
	UserHandler        user.Handler
	TransactionHandler transaction.Handler
}

// Start server
func (s *Server) Start() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(recover.New())
	app.Get("/docs/*", swagger.Handler)

	v1 := app.Group("/api/v1")
	v1.Post("/transactions/income", s.TransactionHandler.SaveIncomeTransaction)
	v1.Post("/transactions/outcome", s.TransactionHandler.SaveOutcomeTransaction)
	v1.Post("/users/register", s.UserHandler.Register)
	v1.Post("/users/login", s.UserHandler.Authenticate)

	err := app.Listen(fmt.Sprintf(":%s", config.Get("PORT")))
	if err != nil {
		fmt.Printf("Server.Start : %v", err)
	}
}
