package model

type (
	RequestPagination struct {
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
		Search   string `json:"search"`
		Sort     string `json:"sort"`
		Filter   string `json:"filter"`
	}

	RequestPaginationSatelliteLog struct {
		RequestPagination
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Status    string `json:"status"`
		Activity  string `json:"activity"`
		Excavator string `json:"excavator"`
	}

	RequestPaginationReportActivity struct {
		RequestPagination
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Excavator string `json:"excavator"`
		Operator  string `json:"operation"`
		Activity  string `json:"activity"`
	}
)
