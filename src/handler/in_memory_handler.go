package handler

import (
	"encoding/json"
	"getir-assignment/src/common"
	"getir-assignment/src/model"
	"getir-assignment/src/service"
	"io"
	"net/http"
)

type InMemoryHandler interface {
	Get(req *http.Request) (int, interface{})
	Post(req *http.Request) (int, interface{})
}

type inMemoryHandler struct {
	svc service.InMemoryService
}

func NewInMemoryHandler(svc service.InMemoryService) InMemoryHandler {
	return &inMemoryHandler{
		svc: svc,
	}
}

// Method: POST
// In-memory storage put key-value endpoint.
// Returns httpStatus and upserted key-value pair.
func (h *inMemoryHandler) Post(req *http.Request) (int, interface{}) {

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	pair := model.Pair{}

	if err := json.Unmarshal(payload, &pair); err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	h.svc.PutKeyValue(pair)

	return http.StatusOK, pair
}

// Method: GET
// In-memory storage get value by key endpoint.
// Returns key-value pair if exists.
func (h *inMemoryHandler) Get(req *http.Request) (int, interface{}) {

	key := req.URL.Query().Get("key")
	if len(key) == 0 {
		errResp := common.GetErrorResponse(common.ErrorInternalServer)
		return errResp.Code, errResp
	}

	result, err := h.svc.GeyValue(key)
	if err != nil {
		errResp := common.GetErrorResponse(err)
		return errResp.Code, errResp
	}

	return 200, result
}
