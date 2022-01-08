package inmem

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockService struct {
	GetMock func(string) (Response, error)
	SetMock func(string, string) (Response, error)
}

func (m mockService) Get(key string) (Response, error) {
	return m.GetMock(key)
}

func (m mockService) Set(key string, value string) (Response, error) {
	return m.SetMock(key, value)
}

func TestController_ServeHTTPValidPost(t *testing.T) {
	mock := mockService{
		SetMock: func(s string, s2 string) (Response, error) {
			return Response{
				Key:   "active-tabs",
				Value: "getir",
				Error: "",
			}, nil
		},
	}

	request := "{\"key\":\"active-tabs\",\"value\":\"getir\"}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/in-memory", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()
	controller := Controller{Repository: mock}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusOK)
	}

	expected := "{\"key\":\"active-tabs\",\"value\":\"getir\"}"
	if rr.Body.String() != expected {
		t.Errorf("returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPMissingField(t *testing.T) {
	mock := mockService{
		SetMock: func(s string, s2 string) (Response, error) {
			return Response{
				Key:   "active-tabs",
				Value: "getir",
				Error: "",
			}, nil
		},
	}

	request := "{\"key\":\"active-tabs\"}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/in-memory", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()
	controller := Controller{Repository: mock}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusBadRequest)
	}

	expected := "{\"key\":\"\",\"value\":\"\",\"error\":\"value field is missing\"}"
	if rr.Body.String() != expected {
		t.Errorf("returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPIntegrationWithService(t *testing.T) {
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
	controller := Controller{Repository: service}

	req, err := http.NewRequest(http.MethodGet, "/in-memory?key=active-tabs", nil)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusOK)
	}

	expected := "{\"key\":\"active-tabs\",\"value\":\"getir\"}"
	if rr.Body.String() != expected {
		t.Errorf("returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}