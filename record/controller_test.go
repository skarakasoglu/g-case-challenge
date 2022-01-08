package record

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var mockData = []Dto{
	{
		Key:        "TAKwGc6Jr4i8Z487",
		CreatedAt:  time.Unix(1485566534, 398000000),
		TotalCount: 310,
	},
	{
		Key:        "LSyjwviN",
		CreatedAt:  time.Unix(1483061467, 831000000),
		TotalCount: 116,
	},
	{
		Key:        "wIFZewQA",
		CreatedAt:  time.Unix(1458343975, 236000000),
		TotalCount: 2863,
	},
}

type mockService struct{
	FetchMock func(options FilterOptions) (Response, error)
}

func (m mockService) Fetch(options FilterOptions) (Response, error) {
	return m.FetchMock(options)
}

func TestController_ServeHTTPValidRequest(t *testing.T) {
	mock := mockService{
		FetchMock: func(options FilterOptions) (Response, error) {
			records := mockData
			return Response{
				Code:    0,
				Message: "Success",
				Records: records,
			}, nil
		},
	}

	request := "{\"startDate\":\"2017-01-27\",\"endDate\":\"2017-01-29\",\"minCount\": 0,\"maxCount\": 1000}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/records", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()
	controller := Controller{Repository: mock}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusOK)
	}

	expected := "{\"code\":0,\"msg\":\"Success\",\"records\":[{\"key\":\"TAKwGc6Jr4i8Z487\",\"createdAt\":\"2017-01-28T01:22:14.398Z\",\"totalCount\":310},{\"key\":\"LSyjwviN\",\"createdAt\":\"2016-12-30T01:31:07.831Z\",\"totalCount\":116},{\"key\":\"wIFZewQA\",\"createdAt\":\"2016-03-18T23:32:55.236Z\",\"totalCount\":2863}]}"
	if rr.Body.String() != expected {
		t.Errorf("returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPNotAllowedMethod(t *testing.T) {
	request := "{\"startDate\": \"2016-01-2x\", \"endDate\": \"2018-02-02\", \"minCount\": 2700, \"maxCount\": 3000}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodGet, "/records", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()

	controller := Controller{}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusMethodNotAllowed)
	}

	expected := "{\"code\":1,\"msg\":\"the method is not allowed for this endpoint.\",\"records\":null}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPByInvalidDate(t *testing.T) {
	request := "{\"startDate\": \"2016-01-2x\", \"endDate\": \"2018-02-02\", \"minCount\": 2700, \"maxCount\": 3000}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/records", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()

	controller := Controller{}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusBadRequest)
	}

	expected := "{\"code\":2,\"msg\":\"parsing time \\\"2016-01-2x\\\" as \\\"2006-01-02\\\": cannot parse \\\"2x\\\" as \\\"02\\\"\",\"records\":null}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPWithAMissingField(t *testing.T) {
	request := "{\"startDate\": \"2016-01-02\", \"minCount\": 2700, \"maxCount\": 3000}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/records", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()

	controller := Controller{}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusBadRequest)
	}

	expected := "{\"code\":2,\"msg\":\"endDate field is missing.\",\"records\":null}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}

func TestController_ServeHTTPIntegrationWithService(t *testing.T) {
	mock := mockDao{
		FindMock: func() ([]Dto, error) {
			return mockData, nil
		},
	}

	request := "{\"startDate\": \"2016-01-01\", \"endDate\": \"2019-01-29\", \"minCount\": 0, \"maxCount\": 3000}"
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, "/records", r)
	if err != nil {
		t.Fatalf("Error on testing: %v", err)
	}

	rr := httptest.NewRecorder()

	service := Service{Dao: mock}
	controller := Controller{Repository: service}
	controller.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned incorrect status code. got: %v, expected: %v", status, http.StatusOK)
	}

	expected := "{\"code\":0,\"msg\":\"Success\",\"records\":[{\"key\":\"TAKwGc6Jr4i8Z487\",\"createdAt\":\"2017-01-28T01:22:14.398Z\",\"totalCount\":310},{\"key\":\"LSyjwviN\",\"createdAt\":\"2016-12-30T01:31:07.831Z\",\"totalCount\":116},{\"key\":\"wIFZewQA\",\"createdAt\":\"2016-03-18T23:32:55.236Z\",\"totalCount\":2863}]}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}