package delivery

import (
	"database/sql"
	"loan/internal/app/investor/repository"
	"loan/internal/app/investor/usecase"
	"loan/internal/pkg/config"

	"github.com/gorilla/mux"
)

type InvestorRegistry struct {
	db      *sql.DB
	handler *InvestorHandler
}

func NewInvestorRegistry(db *sql.DB) *InvestorRegistry {
	repo := repository.NewInvestmentRepository(config.DB)
	usecase := usecase.NewInvestorService(repo)
	handler := NewInvestorHandler(usecase)
	return &InvestorRegistry{db: db, handler: handler}
}

func (h InvestorRegistry) RegisterRoutesTo(r *mux.Router) {
	r.HandleFunc("/loans/{loan_id}/invest", h.handler.InvestInInvestor).Methods("POST")
}
