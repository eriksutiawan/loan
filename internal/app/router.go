package app

import (
	investorweb "loan/internal/app/investor/delivery"
	"loan/internal/app/loan/delivery"
	"loan/internal/pkg/config"

	"github.com/gorilla/mux"
)

type WebHandler struct {
}

func Router(r *mux.Router) {
	db := config.DB

	//investor
	investor := investorweb.NewInvestorRegistry(db)
	investor.RegisterRoutesTo(r)

	//loan
	loan := delivery.NewLoanRegistry(db)
	loan.RegisterRoutesTo(r)
}
