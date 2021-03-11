package transaction

import (
	"github.com/gofiber/fiber/v2"
)

// Handler type definition for passing request to service
type Handler struct {
	TransactionService Service
}

// SaveOutcomeTransaction for handling request and response to save outcome transaction
func (t *Handler) SaveOutcomeTransaction(ctx *fiber.Ctx) error {
	panic("Implement me")
}

// SaveIncomeTransaction for handling request and response to save income transaction
func (t *Handler) SaveIncomeTransaction(ctx *fiber.Ctx) {
	panic("Implement me")
}
