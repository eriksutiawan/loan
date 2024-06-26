package entity

import (
	"database/sql"
)

type Loan struct {
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
