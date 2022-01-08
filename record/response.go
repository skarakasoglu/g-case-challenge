package record

import (
	"strings"
	"time"
)

// Date is declared as a subtype of time.Time to be able to
// override unmarshal function for fields in request JSON with YYYY-MM-dd format.
type Date time.Time

// UnmarshalJSON unmarshals the value if in YYYY-MM-dd format.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

// String is used for logging the type in a meaningful format.
func (d Date) String() string {
	return time.Time(d).String()
}

// RequestPayload represents the request payload.
// The fields are declared as pointer type
// to check whether the fields are supplied in request JSON simply.
type RequestPayload struct{
	StartDate *Date `json:"startDate"`
	EndDate *Date `json:"endDate"`
	MinCount *int `json:"minCount"`
	MaxCount *int `json:"maxCount"`
}

// Response represents the response payload.
type Response struct{
	Code int `json:"code"`
	Message string `json:"msg"`
	Records []Dto `json:"records"`
}