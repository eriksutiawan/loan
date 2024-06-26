package repository

import (
	"context"
	"loan/internal/app/entity"
)

type ILoanRepository interface {
	CreateLoan(ctx context.Context, loan entity.Loan) (*entity.Loan, error)
	UpdateLoanState(ctx context.Context, loanID int, state string) error
	ApproveLoan(ctx context.Context, loanID int, approvalPicture, fieldValidatorID *string, approvalDate string) error
	DisburseLoan(ctx context.Context, loanID int, signedAgreementLetter, fieldOfficerID *string, disbursementDate string) error
	GetLoanDetails(ctx context.Context, loanID int) (*entity.Loan, error)
}
