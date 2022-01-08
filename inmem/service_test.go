package inmem

import (
	"fmt"
	"testing"
)

type mockDao struct{
	GetMock func(string) (Dto, error)
	SetMock func(dto Dto) error
}

func (m mockDao) Get(key string) (Dto, error) {
	return m.GetMock(key)
}

func (m mockDao) Set(dto Dto) error {
	return m.SetMock(dto)
}

func TestService_GetWithExistingKey(t *testing.T) {
	mock := mockDao{
		GetMock: func(s string) (Dto, error) {
			return Dto{
				Key:    "active-tabs",
				Value:  "getir",
				Exists: true,
			}, nil
		},
		SetMock: nil,
	}
	service := Service{Dao: mock}
	got, _ := service.Get("active-tabs")

	expected := Response{
		Key:   "active-tabs",
		Value: "getir",
		Error: "",
	}

	if got.Key != expected.Key {
		t.Errorf("returned incorrect key. got: %v, expected: %v", got.Key, expected.Key)
	}

	if got.Value != expected.Value {
		t.Errorf("returned incorret value. got: %v, expected: %v", got.Value, expected.Value)
	}

	if got.Error != expected.Error {
		t.Errorf("returned incorrect error. got: %v, expected: %v", got.Error, expected.Error)
	}
}

func TestService_GetWithNonExistingKey(t *testing.T) {
	mock := mockDao{
		GetMock: func(s string) (Dto, error) {
			return Dto{
				Key:    "active-tabs",
				Value:  "",
				Exists: false,
			}, nil
		},
		SetMock: nil,
	}
	service := Service{Dao: mock}
	got, _ := service.Get("active-tabs")

	expected := Response{
		Key:   "active-tabs",
		Value: "",
		Error: "key specified does not exist.",
	}

	if got.Key != expected.Key {
		t.Errorf("returned incorrect key. got: %v, expected: %v", got.Key, expected.Key)
	}

	if got.Value != expected.Value {
		t.Errorf("returned incorret value. got: %v, expected: %v", got.Value, expected.Value)
	}

	if got.Error != expected.Error {
		t.Errorf("returned incorrect error. got: %v, expected: %v", got.Error, expected.Error)
	}
}

func TestService_SetSuccess(t *testing.T) {
	mock := mockDao{
		SetMock: func(dto Dto) error {
			return nil
		},
	}
	service := Service{Dao: mock}
	got, _ := service.Set("active-tabs", "getir")

	expected := Response{
		Key:   "active-tabs",
		Value: "getir",
		Error: "",
	}

	if got.Key != expected.Key {
		t.Errorf("returned incorrect key. got: %v, expected: %v", got.Key, expected.Key)
	}

	if got.Value != expected.Value {
		t.Errorf("returned incorret value. got: %v, expected: %v", got.Value, expected.Value)
	}

	if got.Error != expected.Error {
		t.Errorf("returned incorrect error. got: %v, expected: %v", got.Error, expected.Error)
	}
}

func TestService_SetInternalError(t *testing.T) {
	mock := mockDao{
		SetMock: func(dto Dto) error {
			return fmt.Errorf("mock error")
		},
	}
	service := Service{Dao: mock}
	got, _ := service.Set("active-tabs", "getir")

	expected := Response{
		Key:   "active-tabs",
		Value: "",
		Error: "internal server error occurred.",
	}

	if got.Key != expected.Key {
		t.Errorf("returned incorrect key. got: %v, expected: %v", got.Key, expected.Key)
	}

	if got.Value != expected.Value {
		t.Errorf("returned incorret value. got: %v, expected: %v", got.Value, expected.Value)
	}

	if got.Error != expected.Error {
		t.Errorf("returned incorrect error. got: %v, expected: %v", got.Error, expected.Error)
	}
}