package delivery

import (
	"database/sql"
	loanrepo "loan/internal/app/loan/repository"
	loanusecase "loan/internal/app/loan/usecase"
	"loan/internal/pkg/config"

	"github.com/gorilla/mux"
)

type LoanRegistry struct {
	db          *sql.DB
	loanHandler *LoanHandler
}

func NewLoanRegistry(db *sql.DB) *LoanRegistry {
	loanRepo := loanrepo.NewLoanRepository(config.DB)
	loanService := loanusecase.NewLoanService(loanRepo)
	loanHandler := NewLoanHandler(loanService)
	return &LoanRegistry{db: db, loanHandler: loanHandler}
}

func (h LoanRegistry) RegisterRoutesTo(r *mux.Router) {
	r.HandleFunc("/loans", h.loanHandler.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{loan_id}/approve", h.loanHandler.ApproveLoan).Methods("PUT")
	r.HandleFunc("/loans/{loan_id}/disburse", h.loanHandler.DisburseLoan).Methods("PUT")
	r.HandleFunc("/loans/{loan_id}", h.loanHandler.GetLoanDetails).Methods("GET")
}
