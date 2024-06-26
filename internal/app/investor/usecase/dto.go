package usecase

import "github.com/go-playground/validator"

type InvestorCreatorDto struct {
	InvestorID       string  `json:"investor_id" validate:"required"`
	InvestmentAmount float64 `json:"investment_amount" validate:"required,min=0"`
	LoanId           int     `json:"-"`
}

func (dto *InvestorCreatorDto) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
