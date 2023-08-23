package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type requestStruct interface {
	Validate() error
}

func writeJsonResponse(writer http.ResponseWriter, data interface{}) error {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshall error: %s", err)
	}

	_, err = writer.Write(jsonResponse)
	if err != nil {
		return fmt.Errorf("writer.Write error: %s", err)
	}

	return nil
}

func isHTTPMethodAllowed(allowedMethod string, r *http.Request) error {
	if strings.ToUpper(r.Method) != allowedMethod {
		return fmt.Errorf("HTTP method %s is now allowed on route %s, %s is expected", r.Method, r.URL.Path, allowedMethod)
	}
	return nil
}

func populateValidRequest(requestStruct requestStruct, request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, requestStruct)
	if err != nil {
		return err
	}

	err = requestStruct.Validate()
	if err != nil {
		return err
	}

	return nil
}
