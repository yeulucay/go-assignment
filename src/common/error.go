package common

import (
	"errors"
	"getir-assignment/src/model"
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

func GetErrorResponse(err error) model.ResponseModel {
	if errors.Is(err, ErrorBadRequest) {
		return model.ResponseModel{Code: 400, Msg: err.Error()}
	} else if errors.Is(err, ErrorNotFound) {
		return model.ResponseModel{Code: 404, Msg: err.Error()}
	} else if errors.Is(err, ErrorInternalServer) {
		return model.ResponseModel{Code: 500, Msg: err.Error()}
	}

	return model.ResponseModel{Code: 500, Msg: ErrorInternalServer.Error()}
}
