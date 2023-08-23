package db

import (
	"cushon-technical-solution/types"

	"database/sql"
	"fmt"
)

var (
	getCustomerAccountQuery    = "SELECT customerID, investmentFundID, investmentAmount FROM customeraccount WHERE customerID = ? AND investmentFundID = ?"
	updateCustomerAccountQuery = "UPDATE customeraccount SET investmentAmount = ? WHERE customerID = ? AND investmentFundID = ?"
	createCustomerAccountQuery = "INSERT INTO customeraccount (customerID, investmentFund, InvestmentAmount) VALUES (?, ?, ?);"
)

func GetCustomerAccount(customerID int, fundID int) (types.CustomerAccount, error) {
	rows, err := db.Query(getCustomerAccountQuery, customerID, fundID)
	if err != nil {
		return types.CustomerAccount{}, fmt.Errorf("failed to query db: %v", err)
	}

	var customerAccount types.CustomerAccount

	for rows.Next() {
		err := rows.Scan(&customerAccount.CustomerID, &customerAccount.InvestmentFundID, &customerAccount.InvestmentAmount)
		if err == sql.ErrNoRows {
			return types.CustomerAccount{}, err
		} else if err != nil {
			return types.CustomerAccount{}, fmt.Errorf("failed to scan db rows: %v", err)
		}
	}

	err = rows.Close()
	if err != nil {
		return types.CustomerAccount{}, fmt.Errorf("failed to close db rows: %v", err)
	}

	return customerAccount, nil
}

func UpdateCustomerAccount(customerID int, fundID int, amountInvested float64) error {
	_, err := db.Query(updateCustomerAccountQuery, amountInvested, customerID, fundID)
	if err != nil {
		return fmt.Errorf("failed to query db: %v", err)
	}

	return nil
}

func CreateCustomerAccount(customerID int, fundID int, amountInvested float64) error {
	_, err := db.Query(createCustomerAccountQuery, customerID, fundID, amountInvested)
	if err != nil {
		return fmt.Errorf("failed to query db: %v", err)
	}

	return nil
}
