package record

import "time"

// Dto represents the view model of the records to respond the clients in a meaningful format.
type Dto struct{
	Key string `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	TotalCount int `json:"totalCount"`
}
