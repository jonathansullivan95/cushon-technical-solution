package service

import (
	"cushon-technical-solution/db"
	"cushon-technical-solution/types"
)

func GetAvailableFunds() ([]types.Fund, error) {
	return db.GetAvailableFunds()
}
