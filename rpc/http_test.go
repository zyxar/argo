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
	client = NewClient(server)
}

func TestA(t *testing.T) {
	g, err := client.AddUri("http://cran.rstudio.com/bin/macosx/R-3.0.1.pkg")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(g)
	// fmt.Println(client.ForcePause(g))
	// fmt.Println(client.Unpause(g))
	time.Sleep(2 * time.Second)
	client.TellStatus(g)
	fmt.Println(client.Pause(g))
	time.Sleep(1 * time.Second)
	fmt.Println(client.Remove(g))
}
