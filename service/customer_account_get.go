package service

import (
	"cushon-technical-solution/db"
	"cushon-technical-solution/types"
)

func GetCustomerAccount(request types.GetCustomerAccountRequest) (types.CustomerAccount, error) {
	return db.GetCustomerAccount(request.CustomerID, request.InvestmentFundID)
}
