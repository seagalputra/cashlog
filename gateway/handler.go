package gateway

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	. "github.com/seagalputra/cashlog/transaction"
	. "github.com/seagalputra/cashlog/user_account"
)

type GatewayHandler struct {
	TransactionHandler TransactionHandler
	UserAccountHandler UserAccountHandler
}

func (g *GatewayHandler) SaveOutcomeTransaction(ctx *fiber.Ctx) error {
	request := new(CreateOutcomeTransactionRequest)

	if err := ctx.BodyParser(request); err != nil {
		return fmt.Errorf("Failed to parse request body : %v ", err)
	}

	if err := g.TransactionHandler.SaveOutcomeTransaction(*request); err != nil {
		return fmt.Errorf("Failed to save outcome transaction : %v ", err)
	}

	return nil
}
