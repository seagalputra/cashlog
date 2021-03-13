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
	req := new(CreateOutcomeRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := t.TransactionService.CreateOutcome(*req)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated)

	return nil
}

// SaveIncomeTransaction for handling request and response to save income transaction
func (t *Handler) SaveIncomeTransaction(ctx *fiber.Ctx) error {
	req := new(CreateIncomeRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := t.TransactionService.CreateIncome(*req)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated)

	return nil
}
