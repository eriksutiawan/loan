package usecase

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

type CreateLoanDto struct {
	BorrowerID      string  `json:"borrower_id" validate:"required"`
	PrincipalAmount float64 `json:"principal_amount" validate:"required,min=0"`
	Rate            float64 `json:"rate" validate:"required,min=0"`
	ROI             float64 `json:"roi" validate:"required,min=0"`
}

func (dto *CreateLoanDto) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type LoanResponse struct {
	ID                    int          `json:"id"`
	BorrowerID            string       `json:"borrower_id"`
	PrincipalAmount       float64      `json:"principal_amount"`
	Rate                  float64      `json:"rate"`
	ROI                   float64      `json:"roi"`
	State                 string       `json:"state"`
	ApprovalPicture       *string      `json:"approval_picture,omitempty"`
	FieldValidatorID      *string      `json:"field_validator_id,omitempty"`
	ApprovalDate          sql.NullTime `json:"approval_date,omitempty"`
	SignedAgreementLetter *string      `json:"signed_agreement_letter,omitempty"`
	FieldOfficerID        *string      `json:"field_officer_id,omitempty"`
	DisbursementDate      sql.NullTime `json:"disbursement_date,omitempty"`
	AgreementLetterLink   *string      `json:"agreement_letter_link,omitempty"`
	TotalInvestedAmount   float64      `json:"total_invested_amount,omitempty"`
	Investors             []Investor   `json:"investors,omitempty"`
}

type Investor struct {
	InvestorID       string  `json:"investor_id"`
	InvestmentAmount float64 `json:"investment_amount"`
}

type ApprovalDTO struct {
	ApprovalPicture  string `json:"approval_picture" validate:"required"`
	FieldValidatorID string `json:"field_validator_id" validate:"required"`
	ApprovalDate     string `json:"approval_date" validate:"required,customDatetimeFormat"` // Format "2024-06-26 12:00:00"
	LoanId           int    `json:"-"`
}

type DisburseDTO struct {
	SignedAgreementLetter string `json:"signed_agreement_letter" validate:"required"`
	FieldOfficerID        string `json:"field_officer_id" validate:"required"`
	DisbursementDate      string `json:"disbursement_date" validate:"required,customDatetimeFormat"`
	LoanId                int    `json:"-"`
}

func customDatetimeFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02 15:04:05", dateStr)
	return err == nil
}

func (dto *ApprovalDTO) Validate() error {
	validate := validator.New()

	if err := validate.RegisterValidation("customDatetimeFormat", customDatetimeFormat); err != nil {
		return fmt.Errorf("error registering custom validation function: %s", err)
	}

	return validate.Struct(dto)
}

func (dto *DisburseDTO) Validate() error {
	validate := validator.New()

	if err := validate.RegisterValidation("customDatetimeFormat", customDatetimeFormat); err != nil {
		return fmt.Errorf("error registering custom validation function: %s", err)
	}

	return validate.Struct(dto)
}
