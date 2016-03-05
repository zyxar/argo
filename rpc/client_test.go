package rpc

import (
	"fmt"
	"testing"
)

const targetURL = "https://nodejs.org/dist/index.json"

var rpc Protocol

func init() {
	var err error
	rpc, err = New("http://localhost:6800/jsonrpc")
	if err != nil {
		panic(err)
	}
	msg, err := rpc.LaunchAria2cDaemon()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("aria2c", msg.Version, "started!")
	}
	rpc.SetNotifier(&DummyNotifier{})
}

func TestAll(t *testing.T) {
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
	if _, err = rpc.GetURIs(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetFiles(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetPeers(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellWaiting(0, 1); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellStopped(0, 1); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetOption(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetGlobalOption(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetGlobalStat(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.GetSessionInfo(); err != nil {
		t.Error(err)
	}
	if _, err = rpc.Remove(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
	fmt.Println(rpc.ForceShutdown())
}
