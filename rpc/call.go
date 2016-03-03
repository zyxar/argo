package rpc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

var errInvalidParameter = errors.New("invalid parameter")

// Call sends a request of rpc to aria2 daemon
func Call(address, method string, params, reply interface{}) error {
	pay, err := json2.EncodeClientRequest(method, params)
	if err != nil {
		return err
	}
	r, err := http.Post(address, "application/json", bytes.NewReader(pay))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()
	if r.StatusCode != 200 {
		return errors.New(string(body))
	}
	err = json2.DecodeClientResponse(bytes.NewReader(body), &reply)
	return err
}
