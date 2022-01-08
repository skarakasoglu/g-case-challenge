package inmem

// Dao interface is used by a Service to construct a response model
// using data obtained from the database.
type Dao interface{
	Get(string) (Dto, error)
	Set(dto Dto) error
}

// Service uses Dao to access the in-memory database.
// creates responses according to the possible errors.
type Service struct{
	Dao Dao
}

func (s Service) Get(key string) (Response, error) {
	dto, err := s.Dao.Get(key)
	if err != nil {
		return Response{
			Key: key,
			Error: "internal server error occurred.",
		}, err
	}
	
	if !dto.Exists {
		return Response{
			Key:   key,
			Error: "key specified does not exist.",
		}, nil
	}

	resp := Response{
		Key:   key,
		Value: dto.Value,
	}
	return resp, nil
}

func (s Service) Set(key string, value string) (Response, error) {
	err := s.Dao.Set(Dto{
		Key:    key,
		Value:  value,
	})
	if err != nil {
		return Response{
			Key: key,
			Error: "internal server error occurred.",
		}, err
	}

	resp := Response{
		Key:   key,
		Value: value,
	}
	return resp, nil
}