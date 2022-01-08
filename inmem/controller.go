package inmem

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Repository interface{
	Get(string) (Response, error)
	Set(string, string) (Response, error)
}

type Controller struct{
	Repository Repository
}

func (c Controller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// the endpoint always returns JSON response.
	// to notice the client about the content type,
	// it is a great practice to specify in header.
	rw.Header().Add("Content-Type", "application/json")

	switch req.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error on reading the request body: %v", err)
			break
		}

		var payload Request
		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Printf("Error on unmarshalling the request body: %v", err)
			break
		}

		resp, err := c.Repository.Set(payload.Key, payload.Value)
		statusCode := http.StatusOK
		if err != nil {
			log.Printf("Error while setting the value: %v", err)
			statusCode = http.StatusInternalServerError
		}

		c.writeResponse(rw, statusCode, resp)
	case http.MethodGet:
	default:

	}
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