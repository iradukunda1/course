package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iradukunda1/account/render"
)

// Withdraw handler responsible for withdrawing money from an account
func WithDraw(w http.ResponseWriter, r *http.Request) {

	accountNumber := r.FormValue("account_number")
	amount := r.FormValue("amount")

	if accountNumber == "" || amount == "" {
		render.Error(w, fmt.Errorf("missing parameter(s)"), http.StatusBadRequest)
		return
	}

	account, ok := accounts[accountNumber]
	if !ok {
		render.Error(w, fmt.Errorf("account to withdraw from is not found"), http.StatusNotFound)
		return
	}

	if err := account.Withdraw(amount); err != nil {
		render.Error(w, err, http.StatusBadRequest)
		return
	}
	out := RespondData{
		Message: fmt.Sprintf("Withdrawal of %s amount has been withdrawn from account of %s successful.", amount, accountNumber),
	}

	render.Respond(w, out)
}

func (account *BankAccount) Withdraw(in string) error {

	amount, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return err
	}

	if account.Balance-amount < 0 {
		return fmt.Errorf("insufficient balance")
	}

	account.Balance -= amount

	return nil
}
