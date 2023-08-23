package service

import (
	"cushon-technical-solution/db"
	"cushon-technical-solution/types"

	"database/sql"
	"fmt"
)

func UpdateCustomerAccount(request types.UpdateCustomerAccountRequest) error {
	customerAccount, err := db.GetCustomerAccount(request.CustomerID, request.InvestmentFundID)
	if err == sql.ErrNoRows {
		return db.CreateCustomerAccount(request.CustomerID, request.InvestmentFundID, request.InvestmentAmount)
	} else if err != nil {
		return fmt.Errorf("error getting customer account from db: %s", err)
	}

	investmentAmount := customerAccount.InvestmentAmount + request.InvestmentAmount

	err = db.UpdateCustomerAccount(request.CustomerID, request.InvestmentFundID, investmentAmount)
	if err != nil {
		return fmt.Errorf("error updating customer account in db: %s", err)
	}

	return nil
}
