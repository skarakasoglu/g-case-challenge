package inmem

// Dto is used for representing in-memory database key-value pair
// After get operation, if key does not exist, Exists field is set to false
// Otherwise, Exists field will be set to true.
type Dto struct {
	Key string
	Value string
	Exists bool
}
