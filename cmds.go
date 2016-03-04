package main

import (
	"fmt"
	"strconv"
)

var (
	cmds = map[string](func(s ...string) error){
		"launch": func(s ...string) (err error) {
			o, err := rpcc.LaunchAria2cDaemon()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", o)
			return
		},
		"adduri": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddURI(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"addtorrent": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddTorrent(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"addmetalink": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddMetalink(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"remove": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.Remove(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"forceremove": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.ForceRemove(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"pause": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.Pause(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"pauseall": func(s ...string) (err error) {
			ok, err := rpcc.PauseAll()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"forcepause": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.ForcePause(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"forcepauseall": func(s ...string) (err error) {
			ok, err := rpcc.ForcePauseAll()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"unpause": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.Unpause(s[0])
			if err != nil {
				return
			}
			fmt.Printf("gid: %q\n", gid)
			return
		},
		"unpauseall": func(s ...string) (err error) {
			ok, err := rpcc.UnpauseAll()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"tellstatus": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.TellStatus(s[0], s[1:]...)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"geturis": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetURIs(s[0])
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"getfiles": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetFiles(s[0])
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"getpeers": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetPeers(s[0])
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"getservers": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetServers(s[0])
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"tellactive": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.TellActive(s...)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"tellwaiting": func(s ...string) (err error) {
			if len(s) < 2 {
				err = errParameter
				return
			}
			var offset, num int
			if offset, err = strconv.Atoi(s[0]); err != nil {
				return
			}
			if num, err = strconv.Atoi(s[1]); err != nil {
				return
			}
			msg, err := rpcc.TellWaiting(offset, num, s[2:]...)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"tellstopped": func(s ...string) (err error) {
			if len(s) < 2 {
				err = errParameter
				return
			}
			var offset, num int
			if offset, err = strconv.Atoi(s[0]); err != nil {
				return
			}
			if num, err = strconv.Atoi(s[1]); err != nil {
				return
			}
			msg, err := rpcc.TellStopped(offset, num, s[2:]...)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"changeposition": func(s ...string) (err error) {
			if len(s) < 3 {
				err = errParameter
				return
			}
			var pos int
			if pos, err = strconv.Atoi(s[1]); err != nil {
				return
			}
			newp, err := rpcc.ChangePosition(s[0], pos, s[2])
			if err != nil {
				return
			}
			fmt.Printf("newp: %d\n", newp)
			return
		},
		"changeuri": func(s ...string) (err error) {
			err = errNotSupportedCmd
			return
		},
		"option": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetOption(s[0])
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"changeoption": func(s ...string) (err error) {
			err = errNotSupportedCmd
			return
		},
		"globaloption": func(s ...string) (err error) {
			msg, err := rpcc.GetGlobalOption()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"changeglobaloption": func(s ...string) (err error) {
			err = errNotSupportedCmd
			return
		},
		"stat": func(s ...string) (err error) {
			msg, err := rpcc.GetGlobalStat()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"purgeresult": func(s ...string) (err error) {
			ok, err := rpcc.PurgeDownloadResult()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"removeresult": func(s ...string) (err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			ok, err := rpcc.RemoveDownloadResult(s[0])
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"version": func(s ...string) (err error) {
			msg, err := rpcc.GetVersion()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"session": func(s ...string) (err error) {
			msg, err := rpcc.GetSessionInfo()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
		"shutdown": func(s ...string) (err error) {
			ok, err := rpcc.Shutdown()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"forceshutdown": func(s ...string) (err error) {
			ok, err := rpcc.ForceShutdown()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"savesession": func(s ...string) (err error) {
			ok, err := rpcc.SaveSession()
			if err != nil {
				return
			}
			fmt.Println(ok)
			return
		},
		"multicall": func(s ...string) (err error) {
			err = errNotSupportedCmd
			return
		},
		"listmethods": func(s ...string) (err error) {
			msg, err := rpcc.ListMethods()
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msg)
			return
		},
	}
)
