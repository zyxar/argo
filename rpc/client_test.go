package rpc

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

var rpc RPCProto

const (
	serverUri = "http://localhost:6800/jsonrpc"
)

func init() {
	rpc = New(serverUri)
	if err := launchAria2cDaemon(); err != nil {
		panic(err)
	}
}

// func Test0(t *testing.T) {
// 	var reply interface{}
// 	err := Call(serverUri, addUri, []interface{}{[]string{"http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg"}}, &reply)
// 	fmt.Printf("%s %v\n", reply, err)
// }

func TestAll(t *testing.T) {
	defer fmt.Println(rpc.ForceShutdown())
	v, err := rpc.GetVersion()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v["version"])
	g, err := rpc.AddUri("http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(g)
	if _, err = rpc.TellActive(); err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
	if _, err = rpc.TellStatus(g); err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
	if _, err = rpc.PauseAll(); err != nil {
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
	time.Sleep(1. * time.Second)
	fmt.Println("aria2c started!")
	return nil
}
