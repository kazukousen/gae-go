package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Resource interface {
	Get(r *http.Request) (Status, interface{})
}

type Status struct {
	success bool
	code    int
	message string
}

type ResourceBase struct{}

func (ResourceBase) Get(r *http.Request) (Status, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

type apiHeader struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type apiResponse struct {
	Header   apiHeader   `json:"headers"`
	Response interface{} `json:"response"`
}

// ResourceHandler is
func ResourceHandler(Resource Resource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		var status Status
		var data interface{}
		switch r.Method {
		case "GET":
			status, data = Resource.Get(r)
		}

		// Return Response
		var content []byte
		var e error
		if !status.success {
			content, e = json.Marshal(apiResponse{
				Header: apiHeader{Status: "fail", Message: status.message},
			})
		} else {
			content, e = json.Marshal(apiResponse{
				Header:   apiHeader{Status: "success"},
				Response: data,
			})
		}
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)
		w.Write(content)
	}
}

// Fail means API Finished unsuccessfully
func Fail(code int, message string) Status {
	return Status{success: false, code: code, message: message}
}

// FailSimple means API Finished unsuccessfully
func FailSimple(code int) Status {
	return Status{success: false, code: code, message: strconv.Itoa(code) + " " + http.StatusText(code)}
}

// Success means API Finished successfully
func Success(code int) Status {
	return Status{success: true, code: code, message: ""}
}
