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
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "usage: argo {CMD} {PARAMETERS}...\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "available cmds:")
		k := 0
		cmdlist := make([][3]string, len(cmds)/3+1)
		for cmd := range cmds {
			i := k / 3
			j := k % 3
			cmdlist[i][j] = cmd
			k++
		}
		for i := range cmdlist {
			fmt.Fprintf(os.Stderr, "\t%s\r\t\t\t\t%s\r\t\t\t\t\t\t\t\t%s\n", cmdlist[i][0], cmdlist[i][1], cmdlist[i][2])
		}
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
	var err error
	rpcc, err = rpc.New(rpcURI, rpcSecret)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	args := flag.Args()
	if cmd, ok := cmds[args[0]]; ok {
		err = cmd(args[1:]...)
	} else {
		err = errInvalidCmd
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
