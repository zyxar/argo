package rpc

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

const targetURL = "https://nodejs.org/dist/index.json"

var rpc RPCProto

func init() {
	rpc = New("http://localhost:6800/jsonrpc")
	if err := launchAria2cDaemon(); err != nil {
		panic(err)
	}
}

func TestAll(t *testing.T) {
	defer fmt.Println(rpc.ForceShutdown())
	g, err := rpc.AddUri(targetURL)
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
	if o, err = rpc.TellStatus(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.Remove(g); err != nil {
		t.Error(err)
	}
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
}

func launchAria2cDaemon() (err error) {
	if _, err = rpc.GetVersion(); err == nil {
		return
	}
	cmd := exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all")
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Process.Release()
	time.Sleep(1 * time.Second)
	v, err := rpc.GetVersion()
	if err != nil {
		return
	}
	fmt.Println("aria2c", v["version"], "started!")
	return nil
}
