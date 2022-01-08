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
		payload, err := c.parseRequestJSON(req)
		if err != nil {
			log.Printf("Error on parsing request JSON: %v", err)
			return
		}

		resp, err := c.Repository.Set(payload.Key, payload.Value)
		statusCode := http.StatusOK
		if err != nil {
			log.Printf("Error while setting the value: %v", err)
			statusCode = http.StatusInternalServerError
		}

		c.writeResponse(rw, statusCode, resp)
	case http.MethodGet:
		key := req.URL.Query().Get("key")
		resp, err := c.Repository.Get(key)
		statusCode := http.StatusOK
		if err != nil {
			log.Printf("Error while getting the value: %v", err)
			statusCode = http.StatusInternalServerError
		}

		c.writeResponse(rw, statusCode, resp)
	default:

	}
}

// parseRequestJSON reads all request body and unmarshals the JSON to
// Request object.
func (c Controller) parseRequestJSON(r *http.Request) (Request, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Request{}, err
	}

	var payload Request
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return payload, err
	}

	return payload, err
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