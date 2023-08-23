package db

import (
	"cushon-technical-solution/types"

	"fmt"
)

var (
	getAvailableFundsQuery = "SELECT id, name FROM fund"
)

func GetAvailableFunds() ([]types.Fund, error) {
	rows, err := db.Query(getAvailableFundsQuery)
	if err != nil {
		return []types.Fund{}, fmt.Errorf("failed to query db: %v", err)
	}

	var availableFunds []types.Fund

	for rows.Next() {
		var fund types.Fund

		err := rows.Scan(&fund.ID, &fund.Name)
		if err != nil {
			return []types.Fund{}, fmt.Errorf("failed to scan db rows: %v", err)
		}

		availableFunds = append(availableFunds, fund)
	}

	err = rows.Close()
	if err != nil {
		return []types.Fund{}, fmt.Errorf("failed to close db rows: %v", err)
	}

	return availableFunds, nil
}
