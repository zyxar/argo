package rpc

import (
	"fmt"
	"testing"
	"time"
)

var rpc RPCProto

const (
	serverUri = "http://localhost:6800/jsonrpc"
)

func init() {
	rpc = New(serverUri)
}

// func Test0(t *testing.T) {
// 	var reply interface{}
// 	err := Call(serverUri, addUri, []interface{}{[]string{"http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg"}}, &reply)
// 	fmt.Printf("%s %v\n", reply, err)
// }

func TestA(t *testing.T) {
	fmt.Println(rpc.GetVersion())
	g, err := rpc.AddUri("http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(g)
	fmt.Println(rpc.TellActive())
	// fmt.Println(rpc.ForcePause(g))
	// fmt.Println(rpc.Unpause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(rpc.TellStatus(g))
	// time.Sleep(1 * time.Second)
	// fmt.Println(rpc.Pause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(rpc.PauseAll())
	fmt.Println(rpc.Remove(g))
	fmt.Println(rpc.TellActive())
}
