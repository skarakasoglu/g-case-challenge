package record

// Response represents the response payload.
type Response struct{
	Code int `json:"code"`
	Message string `json:"msg"`
	Records []Dto `json:"records"`
}