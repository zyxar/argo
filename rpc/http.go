package rpc

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

func (id *Client) AddUri(uri string, options ...interface{}) (gid string, err error) {
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

func (id *Client) AddTorrent(filename string, options ...interface{}) (gid string, err error) {
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

func (id *Client) AddMetalink(uri string, options ...interface{}) (gid string, err error) {
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

func (id *Client) Remove(gid string) (g string, err error) {
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

func (id *Client) ForceRemove(gid string) (g string, err error) {
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

func (id *Client) Pause(gid string) (g string, err error) {
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

func (id *Client) PauseAll() error {
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

func (id *Client) ForcePause(gid string) (g string, err error) {
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

func (id *Client) Unpause(gid string) (g string, err error) {
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

func (id *Client) UnpauseAll() error {
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

func (id *Client) TellStatus(gid string, keys ...string) error {
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
