// Package record
package record

// Dao interface is used in Service to access data.
type Dao interface{
	Find(options FilterOptions) ([]Dto, error)
}

// Service used by a Controller to interact with a database.
// it implements Repository interface to abstract
// the operations behind the scenes from the Controller and make the code easier to test.
type Service struct{
	Dao Dao
}

// Fetch fetching the records from the Dao by filtering via FilterOptions
// creates a response and returns it.
func (s Service) Fetch(options FilterOptions) (Response, error) {
	records, err := s.Dao.Find(options)
	if err != nil {
		resp := Response{
			Code:    3,
			Message: "internal server error occurred.",
			Records: nil,
		}
		return resp, err
	}

	resp := Response{
		Code:    0,
		Message: "Success",
		Records: records,
	}
	return resp, nil
}
