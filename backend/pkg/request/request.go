package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Error Messages
const dataRetrieveError string = "unable to retrieve request data"

// RetrieveRequestData retrieves the request body data and decodes
// it into the provided data interface
func RetrieveRequestData(r *http.Request, data any) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return errors.New(dataRetrieveError)
	}
	return nil
}
