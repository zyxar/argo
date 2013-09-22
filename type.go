package argo

type request struct {
	Version string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type response struct {
	Version string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type HttpRpc struct {
	uri string
}

func NewHttpRpc(uri string) *HttpRpc {
	return &HttpRpc{uri}
}

type WsRpc struct {
	uri string
}
