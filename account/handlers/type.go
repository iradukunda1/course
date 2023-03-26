package handlers

type BankAccount struct {
	//This represents the account number
	Number string `json:"number"`
	//This represents the account balance
	Balance float64 `json:"balance"`
}

type RespondData struct {
	Message string `json:"message"`
}
