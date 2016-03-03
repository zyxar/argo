package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	cmds = map[string](func(s ...string) ([]byte, error)){
		"launch": func(s ...string) (b []byte, err error) {
			o, err := rpcc.LaunchAria2cDaemon()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(o, "", "  ")
			return
		},
		"adduri": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddURI(s[0])
			if err != nil {
				return
			}
			b = []byte(gid)
			return
		},
		"addtorrent": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddTorrent(s[0])
			if err != nil {
				return
			}
			b = []byte(gid)
			return
		},
		"addmetalink": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			gid, err := rpcc.AddMetalink(s[0])
			if err != nil {
				return
			}
			b = []byte(gid)
			return
		},
		"remove": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.Remove(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"forceremove": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.ForceRemove(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"pause": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.Pause(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"pauseall": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.PauseAll()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"forcepause": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.ForcePause(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"forcepauseall": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.ForcePauseAll()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"unpause": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.Unpause(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"unpauseall": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.UnpauseAll()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"tellstatus": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.TellStatus(s[0], s[1:]...)
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"geturis": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetURIs(s[0])
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"getfiles": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetFiles(s[0])
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"getpeers": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetPeers(s[0])
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"getservers": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetServers(s[0])
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"tellactive": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.TellActive(s...)
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"tellwaiting": func(s ...string) (b []byte, err error) {
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
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"tellstopped": func(s ...string) (b []byte, err error) {
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
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"changeposition": func(s ...string) (b []byte, err error) {
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
			b = []byte(fmt.Sprintf("%d", newp))
			return
		},
		"changeuri": func(s ...string) (b []byte, err error) {
			err = errNotSupportedCmd
			return
		},
		"option": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.GetOption(s[0])
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"changeoption": func(s ...string) (b []byte, err error) {
			err = errNotSupportedCmd
			return
		},
		"globaloption": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.GetGlobalOption()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"changeglobaloption": func(s ...string) (b []byte, err error) {
			err = errNotSupportedCmd
			return
		},
		"stat": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.GetGlobalStat()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"purgeresult": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.PurgeDownloadResult()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"removeresult": func(s ...string) (b []byte, err error) {
			if len(s) == 0 {
				err = errParameter
				return
			}
			msg, err := rpcc.RemoveDownloadResult(s[0])
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"version": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.GetVersion()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"session": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.GetSessionInfo()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
		"shutdown": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.Shutdown()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"forceshutdown": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.ForceShutdown()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"savesession": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.SaveSession()
			if err != nil {
				return
			}
			b = []byte(msg)
			return
		},
		"multicall": func(s ...string) (b []byte, err error) {
			err = errNotSupportedCmd
			return
		},
		"listmethods": func(s ...string) (b []byte, err error) {
			msg, err := rpcc.ListMethods()
			if err != nil {
				return
			}
			b, err = json.MarshalIndent(msg, "", "  ")
			return
		},
	}
)
