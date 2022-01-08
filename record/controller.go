package record

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// FilterOptions the struct is used for
// filtering while fetching records from the database.
type FilterOptions struct{
	StartDate time.Time
	EndDate time.Time
	MinCount int
	MaxCount int
}

// Repository interface provides data needed in Controller.
type Repository interface{
	// Fetch fetches the records from the database according to FilterOptions.
	Fetch(options FilterOptions) (Response, error)
}

// Controller is used for handling "/records" endpoints requests.
// it implements http.Handler interface
// so that the struct can be used as a handler in http.Handle function.
type Controller struct{
	Repository Repository
}

func (c Controller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// the endpoint always returns JSON response.
	// to notice the client about the content type,
	// it is a great practice to specify in header.
	rw.Header().Add("Content-Type", "application/json")

	// for this case, only post method should be handled.
	// the other methods are not allowed.
	switch req.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error on reading the request body: %v", err)
			c.badRequest(rw, "bad request")
			break
		}

		var payload RequestPayload
		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Printf("Error on unmarshalling the request body: %v", err)
			c.badRequest(rw, err.Error())
			break
		}

		log.Printf("/record POST request received. Payload: %+v", payload)

		// should check the validation of the request to make sure the fields are satisfied.
		if !c.validateRequiredFields(rw, payload) {
			break
		}

		statusCode := http.StatusOK
		filterOptions := FilterOptions{
			StartDate: time.Time(*payload.StartDate),
			EndDate:   time.Time(*payload.EndDate),
			MinCount:  *payload.MinCount,
			MaxCount:  *payload.MaxCount,
		}
		resp, err := c.Repository.Fetch(filterOptions)
		if err != nil {
			log.Printf("Error on fetching from the service: %v", err)
			statusCode = http.StatusInternalServerError
		}
		c.writeResponse(rw, statusCode, resp)
	default:
		c.methodNotAllowed(rw)
	}
}

// validateRequiredFields checks whether the fields in request payload are satisfied.
// if not, sends "400 Bad Request" as response.
// Then, it returns if request payload is valid or not.
func (c Controller) validateRequiredFields(rw http.ResponseWriter, payload RequestPayload) bool {
	// since these fields are pointer types, if they are not provided in request JSON
	// after marshalling the JSON to a RequestPayload object, they will be nil.
	// if a field is nil, it means that the field is not provided
	// so the API will return a bad request response with the information of which field is missing.
	if payload.StartDate == nil {
		c.badRequest(rw, "startDate field is missing.")
		return false
	}

	if payload.EndDate == nil {
		c.badRequest(rw, "endDate field is missing.")
		return false
	}

	if payload.MinCount == nil {
		c.badRequest(rw, "minCount field is missing.")
		return false
	}

	if payload.MaxCount == nil {
		c.badRequest(rw, "maxCount field is missing.")
		return false
	}

	return true
}

func (c Controller) badRequest(rw http.ResponseWriter, message string) {
	resp := Response{
		Code:    2,
		Message: message,
	}
	c.writeResponse(rw, http.StatusBadRequest, resp)
}

func (c Controller) methodNotAllowed(rw http.ResponseWriter) {
	resp := Response{
		Code:    1,
		Message: "the method is not allowed for this endpoint.",
	}
	c.writeResponse(rw, http.StatusMethodNotAllowed, resp)
}

// writeResponse converts the response object to byte slice and writes it to response body.
func (c Controller) writeResponse(rw http.ResponseWriter, statusCode int, resp Response) {
	log.Printf("Sending response statusCode: %v, response: %+v", statusCode, resp)

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error on marshalling to JSON: %v", err)
	}

	rw.WriteHeader(statusCode)
	_, err = rw.Write(respBytes)
	if err != nil {
		log.Printf("Error on writing response: %v", err)
	}
}
