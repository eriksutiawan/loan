package usecase

import (
	"context"
	investorrepo "loan/internal/app/investor/repository"
	"loan/internal/pkg/errors"
)

type IInvestorService interface {
	InvestInLoan(ctx context.Context, dto InvestorCreatorDto) *errors.Error
}

type InvestorService struct {
	investmentRepo investorrepo.IInvestmentRepository
}

func NewInvestorService(investmentRepo investorrepo.IInvestmentRepository) IInvestorService {
	return &InvestorService{
		investmentRepo: investmentRepo,
	}
}

func (s *InvestorService) InvestInLoan(ctx context.Context, dto InvestorCreatorDto) *errors.Error {
	if err := dto.Validate(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err := s.investmentRepo.InvestInLoan(ctx, dto.LoanId, dto.InvestorID, dto.InvestmentAmount); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
