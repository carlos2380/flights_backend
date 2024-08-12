package errors

import (
	"encoding/json"
	"net/http"
)

var (
	ErrFetchFlights         = httpError(http.StatusInternalServerError, "failed to fetch flights")
	ErrDecodeFlights        = httpError(http.StatusInternalServerError, "failed to decode flights data")
	ErrMethodNotAllowed     = httpError(http.StatusMethodNotAllowed, "method not allowed")
	ErrUnexpectedStatusCode = httpError(http.StatusInternalServerError, "unexpected status code")
	ErrFetchFlightInfo      = httpError(http.StatusInternalServerError, "failed to fetch flight info")
	ErrDecodeFlightInfo     = httpError(http.StatusInternalServerError, "failed to decode flight info data")
	ErrInvalidFlightID      = httpError(http.StatusBadRequest, "invalid flight ID")
)

func httpError(statusCode int, message string) error {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return e.Message
}

func DetailedError(mainErr, detailedErr error) error {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    mainErr.Error() + ": " + detailedErr.Error(),
	}
}

func WriteJSONError(w http.ResponseWriter, err error, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	_ = json.NewEncoder(w).Encode(
		struct {
			Error string `json:"error"`
		}{Error: err.Error()},
	)
}
