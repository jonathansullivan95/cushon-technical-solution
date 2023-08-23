package api

import (
	"cushon-technical-solution/service"
	"cushon-technical-solution/types"

	"fmt"
	"net/http"
)

func UpdateCustomerAccountInvestmentFundHandler(w http.ResponseWriter, r *http.Request) {
	err := isHTTPMethodAllowed(http.MethodPost, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("isHTTPMethodAllowed error : %v", err)))
		return
	}

	request := types.UpdateCustomerAccountRequest{}
	err = populateValidRequest(&request, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("populateValidRequest error : %v", err)))
		return
	}

	err = service.UpdateCustomerAccount(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to update customers account: %v", err)))
		return
	}

	err = writeJsonResponse(w, "customer account updated succesfully")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to write JSON response: %v", err)))
	}
}
