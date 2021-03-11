package transaction

import "fmt"

type TransactionHandler struct {
	TransactionService TransactionService
}

func (t *TransactionHandler) SaveOutcomeTransaction(request CreateOutcomeTransactionRequest) error {
	err := t.TransactionService.CreateOutcome(request)
	if err != nil {
		return fmt.Errorf("Failed to save outcome transaction : %v ", err)
	}

	return nil
}

func (t *TransactionHandler) SaveIncomeTransaction(request CreateIncomeTransactionRequest) {
	panic("Implement me")
}
