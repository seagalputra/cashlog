package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
