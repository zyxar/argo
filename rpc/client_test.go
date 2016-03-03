package rpc

import (
	"fmt"
	"testing"
)

const targetURL = "https://nodejs.org/dist/index.json"

var rpc Protocol

func init() {
	rpc = New("http://localhost:6800/jsonrpc")
	if msg, err := rpc.LaunchAria2cDaemon(); err != nil {
		panic(err)
	} else {
		fmt.Println("aria2c", msg.Version, "started!")
	}
}

func TestAll(t *testing.T) {
	defer fmt.Println(rpc.ForceShutdown())
	g, err := rpc.AddURI(targetURL)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(g)
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.PauseAll(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellStatus(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.Remove(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
}
