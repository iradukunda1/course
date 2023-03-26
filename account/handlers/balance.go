package handlers

import (
	"fmt"
	"net/http"

	"github.com/iradukunda1/account/render"
)

// Balance handler responsible for getting the balance of an account
func Balance(w http.ResponseWriter, r *http.Request) {

	accountNumber := r.FormValue("account_number")

	if accountNumber == "" {
		render.Error(w, fmt.Errorf("missing parameter: account_number"), http.StatusBadRequest)
		return
	}

	account, ok := accounts[accountNumber]
	if !ok {
		render.Error(w, fmt.Errorf("account not found"), http.StatusNotFound)
		return
	}

	render.Respond(w, account)
}

func (account *BankAccount) GetBalance() float64 {
	return account.Balance
}
