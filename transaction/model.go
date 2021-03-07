package transaction

import (
	. "github.com/seagalputra/cashlog/user_account"
	"math/big"
	"time"
)

type (
	Transaction struct {
		Id              int64
		TransactionId   string
		Title           string
		Amount          big.Float
		TransactionDate time.Time
		Detail          TransactionDetail
		UserAccount     UserAccount
	}

	TransactionDetail struct {
		Id                  int64
		TransactionDetailId string
		Amount              big.Float
		Description         string
		Status              string
	}
)
