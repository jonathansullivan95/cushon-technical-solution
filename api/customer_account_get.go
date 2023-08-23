package api

import (
	"cushon-technical-solution/service"
	"cushon-technical-solution/types"

	"fmt"
	"net/http"
)

func GetCustomerAccountHandler(w http.ResponseWriter, r *http.Request) {
	err := isHTTPMethodAllowed(http.MethodGet, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("isHTTPMethodAllowed error : %v", err)))
	}

	request := types.GetCustomerAccountRequest{}
	err = populateValidRequest(&request, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("populateValidRequest error : %v", err)))
		return
	}

	response, err := service.GetCustomerAccount(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive available funds from db: %v", err)))
	}

	err = writeJsonResponse(w, response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to write JSON response: %v", err)))
	}
}
