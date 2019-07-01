package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/zyxar/argo/rpc"
)

var (
	rpcc               rpc.Client
	rpcSecret          string
	rpcURI             string
	launchLocal        bool
	errParameter       = errors.New("invalid parameter")
	errNotSupportedCmd = errors.New("not supported command")
	errInvalidCmd      = errors.New("invalid command")
)

func init() {
	flag.StringVar(&rpcSecret, "secret", "", "set --rpc-secret for aria2c")
	flag.StringVar(&rpcURI, "uri", "http://localhost:6800/jsonrpc", "set rpc address")
	flag.BoolVar(&launchLocal, "launch", false, "launch local aria2c daemon")
}

func main() {
	flag.Parse()

	if launchLocal {
		if err := LaunchAria2cDaemon(rpcSecret); err != nil {
			fmt.Fprintf(os.Stderr, "launch: %v", err)
			os.Exit(1)
		}
		return
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "usage: argo {CMD} {PARAMETERS}...\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		c := make([]string, 0, len(cmds))
		for cmd := range cmds {
			c = append(c, cmd)
		}
		renderCmdList(os.Stderr, c...)
		os.Exit(1)
	}

	var err error
	rpcc, err = rpc.New(context.Background(), rpcURI, rpcSecret, time.Second, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	defer rpcc.Close()
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

// LaunchAria2cDaemon launchs aria2 daemon to listen for RPC calls, locally.
func LaunchAria2cDaemon(secret string) (err error) {
	args := []string{"--enable-rpc", "--rpc-listen-all"}
	if secret != "" {
		args = append(args, "--rpc-secret="+secret)
	}
	cmd := exec.Command("aria2c", args...)
	if err = cmd.Start(); err != nil {
		return
	}
	return cmd.Process.Release()
}
