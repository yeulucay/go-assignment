package handler

import (
	"encoding/json"
	"getir-assignment/src/common"
	"getir-assignment/src/model"
	"getir-assignment/src/service"
	"io"
	"net/http"
)

type RecordHandler interface {

	// Lists records in database
	// filters records list by model.RecordModel
	List(req *http.Request) (int, interface{})
}

type recordHandler struct {
	svc service.RecordService
}

func NewRecordHandler(svc service.RecordService) RecordHandler {
	return &recordHandler{
		svc: svc,
	}
}

func (h *recordHandler) List(req *http.Request) (int, interface{}) {

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	filter := model.RecordModel{}

	if err := json.Unmarshal(payload, &filter); err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	result, err := h.svc.ListRecordsByFilter(filter)
	if err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	return 200, model.GetSuccessModel(result)
}
