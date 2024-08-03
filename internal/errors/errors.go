package errors

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFetchFlights         = status.New(codes.Internal, "failed to fetch flights").Err()
	ErrDecodeFlights        = status.New(codes.Internal, "failed to decode flights data").Err()
	ErrMethodNotAllowed     = status.New(codes.InvalidArgument, "method not allowed").Err()
	ErrUnexpectedStatusCode = status.New(codes.Internal, "unexpected status code").Err()
)

func DetailedError(mainErr, detailedErr error) error {
	return status.Errorf(codes.Internal, "%v: %v", mainErr, detailedErr)
}

func WriteJSONError(w http.ResponseWriter, err error, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	_ = json.NewEncoder(w).Encode(
		struct {
			Error string `json:"error"`
		}{Error: err.Error()},
	)
}
