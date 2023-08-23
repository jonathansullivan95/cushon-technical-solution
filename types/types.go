package types

import (
	"fmt"
	"strings"
)

type Fund struct {
	ID   int
	Name string
}

type CustomerAccount struct {
	ID               int
	CustomerID       int
	InvestmentFundID int
	InvestmentAmount float64
}

type UpdateCustomerAccountRequest struct {
	CustomerID       int
	InvestmentFundID int
	InvestmentAmount float64
}

func (r UpdateCustomerAccountRequest) Validate() error {
	var errorlist []string
	if r.CustomerID <= 0 {
		errorlist = append(errorlist, invalidError("CustomerID"))
	}
	if r.InvestmentFundID <= 0 {
		errorlist = append(errorlist, invalidError("InvestmentFundID"))
	}
	if r.InvestmentAmount <= 0 {
		errorlist = append(errorlist, invalidError("InvestmentAmount"))
	}

	if len(errorlist) > 0 {
		return fmt.Errorf("validation failed: " + strings.Join(errorlist, ", "))
	}

	return nil
}

type GetCustomerAccountRequest struct {
	CustomerID       int
	InvestmentFundID int
}

func (r GetCustomerAccountRequest) Validate() error {
	var errorlist []string
	if r.CustomerID <= 0 {
		errorlist = append(errorlist, invalidError("CustomerID"))
	}
	if r.InvestmentFundID <= 0 {
		errorlist = append(errorlist, invalidError("InvestmentFundID"))
	}

	if len(errorlist) > 0 {
		return fmt.Errorf("validation failed: " + strings.Join(errorlist, ", "))
	}

	return nil
}

func invalidError(errorField string) string {
	return fmt.Sprintf("%s is invalid", errorField)
}

func zeroOrLessError(errorField string) string {
	return fmt.Sprintf("%s cannot be zero or less", errorField)
}
