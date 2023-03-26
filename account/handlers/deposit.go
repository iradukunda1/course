package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iradukunda1/account/render"
)

// Deposit handler responsible for depositing money to an account
func Deposit(w http.ResponseWriter, r *http.Request) {

	accountNumber := r.FormValue("account_number")
	amount := r.FormValue("amount")

	if accountNumber == "" || amount == "" {
		render.Error(w, fmt.Errorf("missing parameter(s)"), http.StatusBadRequest)
		return
	}

	account, ok := accounts[accountNumber]
	if !ok {
		render.Error(w, fmt.Errorf("account to deposit to is not found"), http.StatusNotFound)
		return
	}

	if err := account.Deposit(amount); err != nil {
		render.Error(w, err, http.StatusBadRequest)
		return
	}

	out := RespondData{
		Message: fmt.Sprintf("Deposit of %s amount has been deposited to %s successful.", amount, accountNumber),
	}

	render.Respond(w, out)

}

// Deposit method responsible for depositing money to an account
func (account *BankAccount) Deposit(in string) error {

	amount, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return err
	}

	if amount < 0 {
		return fmt.Errorf("invalid amount")
	}

	account.Balance += amount
	return nil
}
