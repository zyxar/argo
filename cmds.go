package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/zyxar/argo/rpc"
)

var (
	cmds = map[string](func(s ...string) error){
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
			renderStatusInfo(os.Stdout, msg)
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
			renderURIInfo(os.Stdout, msg...)
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
			renderFileInfo(os.Stdout, msg...)
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
			renderPeerInfo(os.Stdout, msg...)
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
			renderServerInfo(os.Stdout, msg...)
			return
		},
		"tellactive": func(s ...string) (err error) {
			msg, err := rpcc.TellActive(s...)
			if err != nil {
				return
			}
			renderStatusInfo(os.Stdout, msg...)
			return
		},
		"tellwaiting": func(s ...string) (err error) {
			var offset, num int = 0, 10
			if len(s) >= 2 {
				if offset, err = strconv.Atoi(s[0]); err != nil {
					return
				}
				if num, err = strconv.Atoi(s[1]); err != nil {
					return
				}
				s = s[2:]
			}
			msg, err := rpcc.TellWaiting(offset, num, s...)
			if err != nil {
				return
			}
			renderStatusInfo(os.Stdout, msg...)
			return
		},
		"tellstopped": func(s ...string) (err error) {
			var offset, num int = 0, 10
			if len(s) >= 2 {
				if offset, err = strconv.Atoi(s[0]); err != nil {
					return
				}
				if num, err = strconv.Atoi(s[1]); err != nil {
					return
				}
				s = s[2:]
			}
			msg, err := rpcc.TellStopped(offset, num, s...)
			if err != nil {
				return
			}
			renderStatusInfo(os.Stdout, msg...)
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
			renderGlobalStatInfo(os.Stdout, msg)
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
			renderCmdList(os.Stdout, msg...)
			return
		},
	}
)

func renderStatusInfo(w io.Writer, i ...rpc.StatusInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetHeader([]string{"gid", "status", "totalLength", "completedLength", "uploadLength", "bitfield", "downloadSpeed", "uploadSpeed"})
	for _, info := range i {
		tab.Append([]string{
			info.Gid,
			info.Status,
			info.TotalLength,
			info.CompletedLength,
			info.UploadLength,
			info.BitField,
			info.DownloadSpeed,
			info.UploadSpeed,
		})
	}
	tab.Render()
	fmt.Fprintln(w)
}

func renderURIInfo(w io.Writer, i ...rpc.URIInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetHeader([]string{"uri", "status"})
	for _, info := range i {
		tab.Append([]string{
			info.URI,
			info.Status,
		})
	}
	tab.Render()
	fmt.Fprintln(w)
}

func renderFileInfo(w io.Writer, i ...rpc.FileInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetAutoMergeCells(true)
	tab.SetHeader([]string{"index", "path", "length", "completed", "selected", "uri", "status"})
	for _, info := range i {
		for _, uri := range info.URIs {
			tab.Append([]string{
				info.Index,
				info.Path,
				info.Length,
				info.CompletedLength,
				info.Selected,
				uri.URI,
				uri.Status,
			})
		}
	}
	tab.Render()
	fmt.Fprintln(w)
}

func renderPeerInfo(w io.Writer, i ...rpc.PeerInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetHeader([]string{"peerId", "ip", "port", "bitfield", "amChoking", "peerChoking", "downloadSpeed", "uploadSpeed", "seeder"})
	for _, info := range i {
		tab.Append([]string{
			info.PeerId,
			info.IP,
			info.Port,
			info.BitField,
			info.AmChoking,
			info.PeerChoking,
			info.DownloadSpeed,
			info.UploadSpeed,
			info.Seeder,
		})
	}
	tab.Render()
	fmt.Fprintln(w)
}

func renderServerInfo(w io.Writer, i ...rpc.ServerInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetAutoMergeCells(true)
	tab.SetHeader([]string{"index", "uri", "currentUri", "downloadSpeed"})
	for _, info := range i {
		for _, srv := range info.Servers {
			tab.Append([]string{
				info.Index,
				srv.URI,
				srv.CurrentURI,
				srv.DownloadSpeed,
			})
		}
	}
	tab.Render()
	fmt.Fprintln(w)
}

func renderGlobalStatInfo(w io.Writer, info rpc.GlobalStatInfo) {
	tab := tablewriter.NewWriter(w)
	tab.SetHeader([]string{"downloadSpeed", "uploadSpeed", "numActive", "numWaiting", "numStopped", "numStoppedTotal"})
	tab.Append([]string{
		info.DownloadSpeed,
		info.UploadSpeed,
		info.NumActive,
		info.NumWaiting,
		info.NumStopped,
		info.NumStoppedTotal,
	})
	tab.Render()
	fmt.Fprintln(w)
}

func renderCmdList(w io.Writer, cmds ...string) {
	tab := tablewriter.NewWriter(w)
	for i := 0; i < len(cmds)/2; i++ {
		tab.Append([]string{cmds[2*i], "", cmds[2*i+1]})
	}
	if len(cmds)%2 != 0 {
		tab.Append([]string{cmds[len(cmds)-1], "", ""})
	}
	tab.Render()
	fmt.Fprintln(w)
}
