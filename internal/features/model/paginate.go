package model

type (
	RequestPagination struct {
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
		Search   string `json:"search"`
		Sort     string `json:"sort"`
		Filter   string `json:"filter"`
	}

	RequestPaginationTrain struct {
		Page            int    `json:"page"`
		PageSize        int    `json:"page_size"`
		Search          string `json:"search"`
		Sort            string `json:"sort"`
		Filter          string `json:"filter"`
		DepartStationId string `json:"depart_station_id"`
		ArriveStationId string `json:"arrive_station_id"`
	}
)
