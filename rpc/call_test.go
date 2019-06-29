package rpc

import (
	"context"
	"testing"
)

func TestWebsocketCaller(t *testing.T) {
	c, err := newWebsocketCaller(context.Background(), "ws://localhost:6800/jsonrpc")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer c.Close()

	var info VersionInfo
	if err := c.Call(aria2GetVersion, []interface{}{}, &info); err != nil {
		t.Error(err.Error())
	} else {
		println(info.Version)
	}
}
