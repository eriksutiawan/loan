package repository

import (
	"context"
	"database/sql"
)

type InvestmentRepository struct {
	DB *sql.DB
}

func NewInvestmentRepository(db *sql.DB) IInvestmentRepository {
	return &InvestmentRepository{DB: db}
}

func (r *InvestmentRepository) InvestInLoan(ctx context.Context, loanID int, investorID string, investmentAmount float64) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var principalAmount, totalInvested float64
	err = tx.QueryRowContext(ctx, `SELECT principal_amount, COALESCE(SUM(investment_amount), 0) FROM loans LEFT JOIN investments ON loans.id = investments.loan_id WHERE loans.id = $1 GROUP BY loans.id`, loanID).Scan(&principalAmount, &totalInvested)
	if err != nil {
		return err
	}

	if totalInvested+investmentAmount > principalAmount {
		return sql.ErrTxDone
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO investments (loan_id, investor_id, investment_amount) VALUES ($1, $2, $3)`, loanID, investorID, investmentAmount)
	if err != nil {
		return err
	}

	totalInvested += investmentAmount
	if totalInvested == principalAmount {
		_, err = tx.ExecContext(ctx, `UPDATE loans SET state='invested' WHERE id=$1`, loanID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
