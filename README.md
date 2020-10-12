argo/rpc
====

[![Go Report Card](https://goreportcard.com/badge/github.com/zyxar/argo)](https://goreportcard.com/report/github.com/zyxar/argo)
[![GoDoc](https://godoc.org/github.com/zyxar/argo/rpc?status.svg)](https://godoc.org/github.com/zyxar/argo/rpc)
[![Build Status](https://travis-ci.org/zyxar/argo.svg?branch=master)](https://travis-ci.org/zyxar/argo)

aria2 RPC in #Go

## Install

`go get github.com/zyxar/argo/rpc`

## Interface

```go
  AddURI(uris []string, options ...interface{}) (gid string, err error)
  AddTorrent(filename string, options ...interface{}) (gid string, err error)
  AddMetalink(filename string, options ...interface{}) (gid []string, err error)
  Remove(gid string) (msg string, err error)
  ForceRemove(gid string) (msg string, err error)
  Pause(gid string) (msg string, err error)
  PauseAll() (msg string, err error)
  ForcePause(gid string) (msg string, err error)
  ForcePauseAll() (msg string, err error)
  Unpause(gid string) (msg string, err error)
  UnpauseAll() (msg string, err error)
  TellStatus(gid string, keys ...string) (o Option, err error)
  GetUris(gid string) (o Option, err error)
  GetFiles(gid string) (o Option, err error)
  GetPeers(gid string) (o []Option, err error)
  GetServers(gid string) (o []Option, err error)
  TellActive(keys ...string) (o []Option, err error)
  TellWaiting(offset, num int, keys ...string) (o []Option, err error)
  TellStopped(offset, num int, keys ...string) (o []Option, err error)
  ChangePosition(gid string, pos int, how string) (p int, err error)
  ChangeURI(gid string, fileindex int, delUris []string, addUris []string, position ...int) (p []int, err error)
  GetOption(gid string) (o Option, err error)
  ChangeOption(gid string, options Option) (msg string, err error)
  GetGlobalOption() (o Option, err error)
  ChangeGlobalOption(options Option) (msg string, err error)
  GetGlobalStat() (o Option, err error)
  PurgeDowloadResult() (msg string, err error)
  RemoveDownloadResult(gid string) (msg string, err error)
  GetVersion() (o Option, err error)
  GetSessionInfo() (o Option, err error)
  Shutdown() (msg string, err error)
  ForceShutdown() (msg string, err error)
  Multicall(methods []Option) (r []interface{}, err error)
```
