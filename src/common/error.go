package common

import (
	"errors"
)

var (
	ErrorBadRequest     = errors.New("bad request error")
	ErrorNotFound       = errors.New("not found error")
	ErrorInternalServer = errors.New("internal server error")
)

type ErrorResponse struct {
	Msg string `json:"error"`
}

func GetErrorResponse(err error) (int, ErrorResponse) {
	if errors.Is(err, ErrorBadRequest) {
		return 400, ErrorResponse{Msg: err.Error()}
	} else if errors.Is(err, ErrorNotFound) {
		return 404, ErrorResponse{Msg: err.Error()}
	} else if errors.Is(err, ErrorInternalServer) {
		return 500, ErrorResponse{Msg: err.Error()}
	}

	return 500, ErrorResponse{Msg: ErrorInternalServer.Error()}
}
