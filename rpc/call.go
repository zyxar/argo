package rpc

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

// Call sends a request of rpc to aria2 daemon
func Call(address, method string, params, reply interface{}) (err error) {
	pay, err := json2.EncodeClientRequest(method, params)
	if err != nil {
		return
	}
	r, err := http.Post(address, "application/json", bytes.NewReader(pay))
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	r.Body.Close()
	err = json2.DecodeClientResponse(bytes.NewReader(body), &reply)
	return
}
