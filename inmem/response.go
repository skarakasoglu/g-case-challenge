package inmem

// Response represents the response payload.
type Response struct{
	Key string `json:"key"`
	Value string `json:"value"`
	Error string `json:"error,omitempty"`
}
