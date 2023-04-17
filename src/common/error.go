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
	Code int    `json:"code"`
	Msg  string `json:"error"`
}

func GetErrorResponse(err error) ErrorResponse {
	if errors.Is(err, ErrorBadRequest) {
		return ErrorResponse{Code: 400, Msg: err.Error()}
	} else if errors.Is(err, ErrorNotFound) {
		return ErrorResponse{Code: 404, Msg: err.Error()}
	} else if errors.Is(err, ErrorInternalServer) {
		return ErrorResponse{Code: 500, Msg: err.Error()}
	}

	return ErrorResponse{Code: 500, Msg: ErrorInternalServer.Error()}
}
