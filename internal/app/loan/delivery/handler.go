package delivery

import (
	"context"
	"encoding/json"
	"loan/internal/app/loan/usecase"
	"loan/internal/pkg/errors"
	"loan/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LoanHandler struct {
	loanService usecase.ILoanService
}

func NewLoanHandler(loanService usecase.ILoanService) *LoanHandler {
	return &LoanHandler{loanService: loanService}
}

func (h *LoanHandler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	var loan usecase.CreateLoanDto
	err := json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	svcErr := h.loanService.CreateLoan(ctx, loan)
	if svcErr != nil {
		response.RespondWithJSON(w, svcErr.Code, svcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loan)
}

func (h *LoanHandler) ApproveLoan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanID := vars["loan_id"]
	if loanID == "" {
		badErr := errors.NewBadRequestError("loand_id required")
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	loanId, err := strconv.Atoi(loanID)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	var loan usecase.ApprovalDTO
	err = json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	loan.LoanId = loanId
	ctx := context.Background()
	aproveErr := h.loanService.ApproveLoan(ctx, loan)
	if aproveErr != nil {
		response.RespondWithJSON(w, aproveErr.Code, aproveErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Loan approved"})
}

func (h *LoanHandler) DisburseLoan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanID := vars["loan_id"]
	if loanID == "" {
		badErr := errors.NewBadRequestError("loand_id required")
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	loanId, err := strconv.Atoi(loanID)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	var loan usecase.DisburseDTO
	err = json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	ctx := context.Background()
	loan.LoanId = loanId
	disburseErr := h.loanService.DisburseLoan(ctx, loan)
	if disburseErr != nil {
		response.RespondWithJSON(w, disburseErr.Code, disburseErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Loan disbursed"})
}

func (h *LoanHandler) GetLoanDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanID := vars["loan_id"]
	if loanID == "" {
		badErr := errors.NewBadRequestError("loand_id required")
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	loanId, err := strconv.Atoi(loanID)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	ctx := context.Background()
	loan, loanErr := h.loanService.GetLoanDetails(ctx, loanId)
	if loanErr != nil {
		response.RespondWithJSON(w, loanErr.Code, loanErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loan)
}
