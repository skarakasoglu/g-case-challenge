package inmem

// Request represents the request payload.
type Request struct{
	Key *string `json:"key"`
	Value *string `json:"value"`
}
