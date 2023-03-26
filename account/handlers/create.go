package handlers

import (
	"fmt"
	"net/http"

	"github.com/iradukunda1/account/render"
)

// accounts that will store all the accounts information with the account number as the key
var accounts = make(map[string]*BankAccount)

// Create handler responsible for creating a new account
func Create(w http.ResponseWriter, r *http.Request) {

	accountNumber := r.FormValue("account_number")

	if accountNumber == "" {
		render.Error(w, fmt.Errorf("missing parameter: account_number"), http.StatusBadRequest)
		return
	}

	if _, ok := accounts[accountNumber]; ok {
		render.Error(w, fmt.Errorf("account already exists"), http.StatusBadRequest)
		return
	}

	accounts[accountNumber] = &BankAccount{accountNumber, 0}

	render.Respond(w, RespondData{fmt.Sprintf("Account %s has been created", accountNumber)})
}
