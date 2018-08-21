package rpc

import (
	"bytes"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
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
	pay, err := json2.EncodeClientRequest(method, params)
	if err != nil {
		return
	}
	r, err := http.Post(string(h), "application/json", bytes.NewReader(pay))
	if err != nil {
		return
	}
	err = json2.DecodeClientResponse(r.Body, &reply)
	r.Body.Close()
	return
}
