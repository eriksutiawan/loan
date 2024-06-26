package repository

import (
	"context"
)

type IInvestmentRepository interface {
	InvestInLoan(ctx context.Context, loanID int, investorID string, investmentAmount float64) error
}
