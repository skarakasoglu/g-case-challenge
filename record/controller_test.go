package record

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

	expected := "{\n    \"code\": 2,\n    \"msg\": \"parsing time \\\"2016-01-2x\\\" as \\\"2006-01-02\\\": cannot parse \\\"2x\\\" as \\\"02\\\"\",\n    \"records\": null\n}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned incorrect response body. got: %v, expected: %v", rr.Body.String(), expected)
	}
}
