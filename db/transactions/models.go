// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package transactions

import (
	"database/sql"

	"github.com/jackc/pgtype"
)

type Balance struct {
	UserID           int32
	CurrencyID       int32
	BalanceAmount    int64
	TransactionCount int64
}

type Transaction struct {
	TransactionID     string
	UserID            int32
	CurrencyID        int32
	TransactionAmount pgtype.Numeric
	TransactionDate   sql.NullTime
}
