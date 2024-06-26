package usecase

import (
	"context"
	"database/sql"
	"loan/internal/app/entity"
	loanrepo "loan/internal/app/loan/repository"
	pkgError "loan/internal/pkg/errors"
)

type ILoanService interface {
	CreateLoan(ctx context.Context, dto CreateLoanDto) *pkgError.Error
	ApproveLoan(ctx context.Context, dto ApprovalDTO) *pkgError.Error
	DisburseLoan(ctx context.Context, dto DisburseDTO) *pkgError.Error
	GetLoanDetails(ctx context.Context, loanID int) (*LoanResponse, *pkgError.Error)
}

type LoanService struct {
	loanRepo loanrepo.ILoanRepository
}

func NewLoanService(loanRepo loanrepo.ILoanRepository) ILoanService {
	return &LoanService{
		loanRepo: loanRepo,
	}
}

func (s *LoanService) CreateLoan(ctx context.Context, dto CreateLoanDto) *pkgError.Error {
	if err := dto.Validate(); err != nil {
		return pkgError.NewBadRequestError(err.Error())
	}
	_, err := s.loanRepo.CreateLoan(ctx, entity.Loan{
		BorrowerID:      dto.BorrowerID,
		PrincipalAmount: dto.PrincipalAmount,
		ROI:             dto.ROI,
		Rate:            dto.Rate,
	})

	if err != nil {
		return pkgError.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *LoanService) ApproveLoan(ctx context.Context, dto ApprovalDTO) *pkgError.Error {
	if err := dto.Validate(); err != nil {
		return pkgError.NewBadRequestError(err.Error())
	}

	err := s.loanRepo.ApproveLoan(ctx, dto.LoanId, &dto.ApprovalPicture, &dto.FieldValidatorID, dto.ApprovalDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkgError.NewEntityNotFound(err.Error())
		}
		return pkgError.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *LoanService) DisburseLoan(ctx context.Context, dto DisburseDTO) *pkgError.Error {
	if err := dto.Validate(); err != nil {
		return pkgError.NewBadRequestError(err.Error())
	}

	err := s.loanRepo.DisburseLoan(ctx, dto.LoanId, &dto.SignedAgreementLetter, &dto.FieldOfficerID, dto.DisbursementDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkgError.NewEntityNotFound(err.Error())
		}
		return pkgError.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *LoanService) GetLoanDetails(ctx context.Context, loanID int) (*LoanResponse, *pkgError.Error) {
	loan, err := s.loanRepo.GetLoanDetails(ctx, loanID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkgError.NewEntityNotFound(err.Error())
		}
		return nil, pkgError.NewInternalServerError(err.Error())
	}

	var investors []Investor
	for _, investor := range loan.Investors {
		investors = append(investors, Investor{
			InvestorID:       investor.InvestorID,
			InvestmentAmount: investor.InvestmentAmount,
		})
	}

	return &LoanResponse{
		ID:                    loan.ID,
		BorrowerID:            loan.BorrowerID,
		PrincipalAmount:       loan.PrincipalAmount,
		Rate:                  loan.Rate,
		ROI:                   loan.ROI,
		State:                 loan.State,
		ApprovalPicture:       loan.ApprovalPicture,
		FieldValidatorID:      loan.FieldValidatorID,
		ApprovalDate:          loan.ApprovalDate,
		SignedAgreementLetter: loan.SignedAgreementLetter,
		FieldOfficerID:        loan.FieldOfficerID,
		DisbursementDate:      loan.DisbursementDate,
		AgreementLetterLink:   loan.AgreementLetterLink,
		TotalInvestedAmount:   loan.TotalInvestedAmount,
		Investors:             investors,
	}, nil
}
