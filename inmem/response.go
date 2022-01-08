// Package inmem manages operations between in-memory database.
package inmem

// Response represents the response payload.
type Response struct{
	Key string `json:"key"`
	Value string `json:"value"`
}
