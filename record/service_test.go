package record

import (
	"fmt"
	"testing"
)

type mockDao struct{
	FindMock func() ([]Dto, error)
}

func (m mockDao) Find(options FilterOptions) ([]Dto, error) {
	return m.FindMock()
}

func TestService_FetchSuccess(t *testing.T) {
	mock := mockDao{
		FindMock: func() ([]Dto, error) {
			return mockData, nil
		},
	}

	service := Service{Dao: mock}
	got, _ := service.Fetch(FilterOptions{})

	expected := Response{
		Code:    0,
		Message: "Success",
		Records: mockData,
	}

	if got.Code != expected.Code {
		t.Errorf("returned incorrect response code. got: %v, expected: %v", got.Code, expected.Code)
	}

	if got.Message != expected.Message {
		t.Errorf("returned incorrect response message. got: %v, expected: %v", got.Message, expected.Message)
	}
}

func TestService_FetchInternalError(t *testing.T) {
	mock := mockDao{
		FindMock: func() ([]Dto, error) {
			return mockData, fmt.Errorf("error")
		},
	}

	service := Service{Dao: mock}
	got, _ := service.Fetch(FilterOptions{})

	expected := Response{
		Code:    3,
		Message: "internal server error occurred.",
		Records: mockData,
	}

	if got.Code != expected.Code {
		t.Errorf("returned incorrect response code. got: %v, expected: %v", got.Code, expected.Code)
	}

	if got.Message != expected.Message {
		t.Errorf("returned incorrect response message. got: %v, expected: %v", got.Message, expected.Message)
	}
}
