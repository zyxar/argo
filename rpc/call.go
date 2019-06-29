package rpc

import (
	"net/http"
)

type caller interface {
	// Call sends a request of rpc to aria2 daemon
	Call(method string, params, reply interface{}) (err error)
}

type httpCaller string

func newHTTPCaller(uri string) caller {
	return httpCaller(uri)
}

func (h httpCaller) Call(method string, params, reply interface{}) (err error) {
	payload, err := EncodeClientRequest(method, params)
	if err != nil {
		return
	}
	r, err := http.Post(string(h), "application/json", payload)
	if err != nil {
		return
	}
	err = DecodeClientResponse(r.Body, &reply)
	r.Body.Close()
	return
}
