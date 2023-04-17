package model

type ResponseModel struct {
	Code    int         `json:"code"`
	Msg     string      `json:"error"`
	Records interface{} `json:"records,omitempty"`
}

func GetSuccessModel(records interface{}) ResponseModel {
	return ResponseModel{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}
}
