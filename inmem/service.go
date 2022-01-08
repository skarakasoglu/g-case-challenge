package inmem

type Service struct{
	dao Dao
}

func (s Service) Get(key string) (Response, error) {
	val, err := s.dao.Get(key)
	if err != nil {
		return Response{}, err
	}

	resp := Response{
		Key:   key,
		Value: val,
	}
	return resp, nil
}

func (s Service) Set(key string, value string) (Response, error) {
	err := s.dao.Set(key, value)
	if err != nil {
		return Response{}, err
	}

	resp := Response{
		Key:   key,
		Value: value,
	}
	return resp, nil
}