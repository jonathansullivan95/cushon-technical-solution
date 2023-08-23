package api

import (
	"cushon-technical-solution/service"
	"fmt"
	"net/http"
)

func GetAvailableFundsHandler(w http.ResponseWriter, r *http.Request) {
	err := isHTTPMethodAllowed(http.MethodGet, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("isHTTPMethodAllowed error : %v", err)))
	}

	response, err := service.GetAvailableFunds()
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
