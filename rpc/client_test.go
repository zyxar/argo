package rpc

import (
	"fmt"
	"testing"
	"time"
)

var client *Client

const (
	server = "http://localhost:6800/jsonrpc"
)

func init() {
	client = New(server)
}

// func Test0(t *testing.T) {
// 	var reply interface{}
// 	err := Call(server, addUri, []interface{}{[]string{"http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg"}}, &reply)
// 	fmt.Printf("%s %v\n", reply, err)
// }

func TestA(t *testing.T) {
	fmt.Println(client.GetVersion())
	g, err := client.AddUri("http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(g)
	fmt.Println(client.TellActive())
	// fmt.Println(client.ForcePause(g))
	// fmt.Println(client.Unpause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(client.TellStatus(g))
	// time.Sleep(1 * time.Second)
	// fmt.Println(client.Pause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(client.PauseAll())
	fmt.Println(client.Remove(g))
	fmt.Println(client.TellActive())
}
