package argo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *request) Pack() []byte {
	b, _ := json.Marshal(r)
	return b
}

func (id *HttpRpc) AddUri(uri string, options ...interface{}) (gid string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.addUri",
		Params:  []interface{}{[]string{uri}},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	gid = r.Result
	return
}

func (id *HttpRpc) AddTorrent(filename string, options ...interface{}) (gid string, err error) {
	co, err := ioutil.ReadFile(filename) // to base64
	if err != nil {
		return
	}
	file := base64.StdEncoding.EncodeToString(co)
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.addTorrent",
		Params:  []interface{}{string(file)},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	gid = r.Result
	return
}

func (id *HttpRpc) AddMetalink(uri string, options ...interface{}) (gid string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.addMetalink",
		Params:  []interface{}{uri},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	gid = r.Result
	return
}

func (id *HttpRpc) Remove(gid string) (g string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.remove",
		Params:  []interface{}{gid},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	g = r.Result
	return
}

func (id *HttpRpc) ForceRemove(gid string) (g string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.forceRemove",
		Params:  []interface{}{gid},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	g = r.Result
	return
}

func (id *HttpRpc) Pause(gid string) (g string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.pause",
		Params:  []interface{}{gid},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	g = r.Result
	return
}

func (id *HttpRpc) PauseAll() error {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.pauseAll",
		Params:  nil,
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
	return err
}

func (id *HttpRpc) ForcePause(gid string) (g string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.forcePause",
		Params:  []interface{}{gid},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	g = r.Result
	return
}

func (id *HttpRpc) Unpause(gid string) (g string, err error) {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.unpause",
		Params:  []interface{}{gid},
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r response
	if err = json.Unmarshal(body, &r); err != nil {
		return
	}
	if r.Error.Code != 0 {
		fmt.Printf("%s\n", r.Error.Message)
		return
	}
	g = r.Result
	return
}

func (id *HttpRpc) UnpauseAll() error {
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.unpauseAll",
		Params:  nil,
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
	return err
}

func (id *HttpRpc) TellStatus(gid string, keys ...string) error {
	pay := make([]interface{}, 1, len(keys)+1)
	pay[0] = gid
	for i, _ := range keys {
		pay = append(pay, keys[i])
	}
	req := &request{
		Version: "2.0",
		Id:      "qwer",
		Method:  "aria2.tellStatus",
		Params:  pay,
	}
	resp, err := http.Post(id.uri, "application/json", bytes.NewReader(req.Pack()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
	return err
}
