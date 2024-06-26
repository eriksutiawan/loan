package repository

import (
	"context"
	"database/sql"
	"loan/internal/app/entity"
)

type LoanRepository struct {
	DB *sql.DB
}

func NewLoanRepository(db *sql.DB) ILoanRepository {
	return &LoanRepository{DB: db}
}

func (r *LoanRepository) CreateLoan(ctx context.Context, loan entity.Loan) (*entity.Loan, error) {
	query := `INSERT INTO loans (borrower_id, principal_amount, rate, roi, state) VALUES ($1, $2, $3, $4, 'proposed') RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, loan.BorrowerID, loan.PrincipalAmount, loan.Rate, loan.ROI).Scan(&loan.ID)
	return &loan, err
}

func (r *LoanRepository) UpdateLoanState(ctx context.Context, loanID int, state string) error {
	query := `UPDATE loans SET state=$1 WHERE id=$2`
	_, err := r.DB.ExecContext(ctx, query, state, loanID)
	return err
}

func (r *LoanRepository) ApproveLoan(ctx context.Context, loanID int, approvalPicture, fieldValidatorID *string, approvalDate string) error {
	query := `UPDATE loans SET state='approved', approval_picture=$1, field_validator_id=$2, approval_date=$3 WHERE id=$4 AND state='proposed'`
	res, err := r.DB.ExecContext(ctx, query, approvalPicture, fieldValidatorID, approvalDate, loanID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *LoanRepository) DisburseLoan(ctx context.Context, loanID int, signedAgreementLetter *string, fieldOfficerID *string, disbursementDate string) error {
	query := `UPDATE loans SET state='disbursed', signed_agreement_letter=$1, field_officer_id=$2, disbursement_date=$3 WHERE id=$4 AND state='invested'`
	res, err := r.DB.ExecContext(ctx, query, signedAgreementLetter, fieldOfficerID, disbursementDate, loanID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *LoanRepository) GetLoanDetails(ctx context.Context, loanID int) (*entity.Loan, error) {
	var loan entity.Loan
	err := r.DB.QueryRowContext(ctx, `SELECT id, borrower_id, principal_amount, rate, roi, state, approval_picture, field_validator_id, approval_date, signed_agreement_letter, field_officer_id, disbursement_date, agreement_letter_link FROM loans WHERE id=$1`, loanID).Scan(&loan.ID, &loan.BorrowerID, &loan.PrincipalAmount, &loan.Rate, &loan.ROI, &loan.State, &loan.ApprovalPicture, &loan.FieldValidatorID, &loan.ApprovalDate, &loan.SignedAgreementLetter, &loan.FieldOfficerID, &loan.DisbursementDate, &loan.AgreementLetterLink)
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, `SELECT investor_id, investment_amount FROM investments WHERE loan_id=$1`, loanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var investor entity.Investor
		if err := rows.Scan(&investor.InvestorID, &investor.InvestmentAmount); err != nil {
			return nil, err
		}
		loan.Investors = append(loan.Investors, investor)
	}

	return &loan, nil
}
