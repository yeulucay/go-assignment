package model

type RecordModel struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type RecordResult struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int64  `json:"totalCount"`
}
