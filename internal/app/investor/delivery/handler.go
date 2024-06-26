package delivery

import (
	"context"
	"encoding/json"
	"loan/internal/app/investor/usecase"
	"loan/internal/pkg/errors"
	"loan/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvestorHandler struct {
	investorService usecase.IInvestorService
}

func NewInvestorHandler(investorService usecase.IInvestorService) *InvestorHandler {
	return &InvestorHandler{investorService: investorService}
}

func (h *InvestorHandler) InvestInInvestor(w http.ResponseWriter, r *http.Request) {
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

	var investor usecase.InvestorCreatorDto
	err = json.NewDecoder(r.Body).Decode(&investor)
	if err != nil {
		badErr := errors.NewBadRequestError(err.Error())
		response.RespondWithJSON(w, badErr.Code, badErr)
		return
	}

	investor.LoanId = loanId
	ctx := context.Background()
	investErr := h.investorService.InvestInLoan(ctx, investor)
	if investErr != nil {
		response.RespondWithJSON(w, investErr.Code, investErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Investment successful"})
}
