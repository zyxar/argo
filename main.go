package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/zyxar/argo/rpc"
)

var (
	rpcc               rpc.Protocol
	rpcSecret          string
	rpcURI             string
	errParameter       = errors.New("invalid parameter")
	errNotSupportedCmd = errors.New("not supported command")
	errInvalidCmd      = errors.New("invalid command")
)

func init() {
	flag.StringVar(&rpcSecret, "secret", "", "set --rpc-secret for aria2c")
	flag.StringVar(&rpcURI, "uri", "http://localhost:6800/jsonrpc", "set rpc address")
}

func main() {
	flag.Parse()
	rpcc = rpc.New(rpcURI, rpcSecret)
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "usage: argo {CMD} {PARAMETERS}...\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	args := flag.Args()
	var err error
	if cmd, ok := cmds[args[0]]; ok {
		err = cmd(args[1:]...)
	} else {
		err = errInvalidCmd
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
