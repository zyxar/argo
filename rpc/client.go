package rpc

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os/exec"
	"sync"
	"time"
)

// Option is a container for specifying Call parameters and returning results
type Option map[string]interface{}

// Protocol is a set of rpc methods that aria2 daemon supports
type Protocol interface {
	LaunchAria2cDaemon() (info VersionInfo, err error)
	AddURI(uri string, options ...interface{}) (gid string, err error)
	AddTorrent(filename string, options ...interface{}) (gid string, err error)
	AddMetalink(uri string, options ...interface{}) (gid string, err error)
	Remove(gid string) (g string, err error)
	ForceRemove(gid string) (g string, err error)
	Pause(gid string) (g string, err error)
	PauseAll() (ok string, err error)
	ForcePause(gid string) (g string, err error)
	ForcePauseAll() (ok string, err error)
	Unpause(gid string) (g string, err error)
	UnpauseAll() (ok string, err error)
	TellStatus(gid string, keys ...string) (info StatusInfo, err error)
	GetURIs(gid string) (infos []URIInfo, err error)
	GetFiles(gid string) (infos []FileInfo, err error)
	GetPeers(gid string) (infos []PeerInfo, err error)
	GetServers(gid string) (infos []ServerInfo, err error)
	TellActive(keys ...string) (infos []StatusInfo, err error)
	TellWaiting(offset, num int, keys ...string) (infos []StatusInfo, err error)
	TellStopped(offset, num int, keys ...string) (infos []StatusInfo, err error)
	ChangePosition(gid string, pos int, how string) (p int, err error)
	ChangeURI(gid string, fileindex int, delUris []string, addUris []string, position ...int) (p []int, err error)
	GetOption(gid string) (m Option, err error)
	ChangeOption(gid string, option Option) (ok string, err error)
	GetGlobalOption() (m Option, err error)
	ChangeGlobalOption(options Option) (ok string, err error)
	GetGlobalStat() (info GlobalStatInfo, err error)
	PurgeDownloadResult() (ok string, err error)
	RemoveDownloadResult(gid string) (ok string, err error)
	GetVersion() (info VersionInfo, err error)
	GetSessionInfo() (info SessionInfo, err error)
	Shutdown() (ok string, err error)
	ForceShutdown() (ok string, err error)
	SaveSession() (ok string, err error)
	Multicall(methods []Method) (r []interface{}, err error)
	ListMethods() (methods []string, err error)
}

type client struct {
	mutex sync.Mutex
	uri   string
	token string
}

var errInvalidParameter = errors.New("invalid parameter")

// New returns an instance of Protocol
func New(s ...string) Protocol {
	switch len(s) {
	case 0:
		return nil
	case 1:
		return &client{uri: s[0]}
	}
	return &client{uri: s[0], token: s[1]}
}

func (id *client) lock() {
	id.mutex.Lock()
}

func (id *client) unlock() {
	id.mutex.Unlock()
}

// LaunchAria2cDaemon launchs aria2 daemon to listen for RPC calls, locally.
func (id *client) LaunchAria2cDaemon() (info VersionInfo, err error) {
	if info, err = id.GetVersion(); err == nil {
		return
	}
	args := []string{"--enable-rpc", "--rpc-listen-all"}
	if id.token != "" {
		args = append(args, "--rpc-secret="+id.token)
	}
	cmd := exec.Command("aria2c", args...)
	if err = cmd.Start(); err != nil {
		return
	}
	cmd.Process.Release()
	time.Sleep(1 * time.Second)
	info, err = id.GetVersion()
	return
}

// `aria2.addUri([secret, ]uris[, options[, position]])`
// This method adds a new download. uris is an array of HTTP/FTP/SFTP/BitTorrent URIs (strings) pointing to the same resource.
// If you mix URIs pointing to different resources, then the download may fail or be corrupted without aria2 complaining.
// When adding BitTorrent Magnet URIs, uris must have only one element and it should be BitTorrent Magnet URI.
// options is a struct and its members are pairs of option name and value.
// If position is given, it must be an integer starting from 0.
// The new download will be inserted at position in the waiting queue.
// If position is omitted or position is larger than the current size of the queue, the new download is appended to the end of the queue.
// This method returns the GID of the newly registered download.
func (id *client) AddURI(uri string, options ...interface{}) (gid string, err error) {
	params := make([]interface{}, 1, 2)
	params[0] = []string{uri}
	if options != nil {
		params = append(params, options...)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2AddURI, params, &gid)
	id.unlock()
	return
}

// `aria2.addTorrent([secret, ]torrent[, uris[, options[, position]]])`
// This method adds a BitTorrent download by uploading a ".torrent" file.
// If you want to add a BitTorrent Magnet URI, use the aria2.addUri() method instead.
// torrent must be a base64-encoded string containing the contents of the ".torrent" file.
// uris is an array of URIs (string). uris is used for Web-seeding.
// For single file torrents, the URI can be a complete URI pointing to the resource; if URI ends with /, name in torrent file is added.
// For multi-file torrents, name and path in torrent are added to form a URI for each file. options is a struct and its members are pairs of option name and value.
// If position is given, it must be an integer starting from 0.
// The new download will be inserted at position in the waiting queue.
// If position is omitted or position is larger than the current size of the queue, the new download is appended to the end of the queue.
// This method returns the GID of the newly registered download.
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named as the hex string of SHA-1 hash of data plus ".torrent" in the directory specified by --dir option.
// E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.torrent.
// If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (id *client) AddTorrent(filename string, options ...interface{}) (gid string, err error) {
	co, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	file := base64.StdEncoding.EncodeToString(co)
	params := make([]interface{}, 1, 2)
	params[0] = string(file)
	if options != nil {
		params = append(params, options...)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2AddTorrent, params, &gid)
	id.unlock()
	return
}

// `aria2.addMetalink([secret, ]metalink[, options[, position]])`
// This method adds a Metalink download by uploading a ".metalink" file.
// metalink is a base64-encoded string which contains the contents of the ".metalink" file.
// options is a struct and its members are pairs of option name and value.
// If position is given, it must be an integer starting from 0.
// The new download will be inserted at position in the waiting queue.
// If position is omitted or position is larger than the current size of the queue, the new download is appended to the end of the queue.
// This method returns an array of GIDs of newly registered downloads.
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named hex string of SHA-1 hash of data plus ".metalink" in the directory specified by --dir option.
// E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.metalink.
// If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (id *client) AddMetalink(uri string, options ...interface{}) (gid string, err error) {
	params := make([]interface{}, 1, 2)
	params[0] = uri
	if options != nil {
		params = append(params, options...)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2AddMetalink, params, &gid)
	id.unlock()
	return
}

// `aria2.remove([secret, ]gid)`
// This method removes the download denoted by gid (string).
// If the specified download is in progress, it is first stopped.
// The status of the removed download becomes removed.
// This method returns GID of removed download.
func (id *client) Remove(gid string) (g string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2Remove, params, &g)
	id.unlock()
	return
}

// `aria2.forceRemove([secret, ]gid)`
// This method removes the download denoted by gid.
// This method behaves just like aria2.remove() except that this method removes the download without performing any actions which take time, such as contacting BitTorrent trackers to unregister the download first.
func (id *client) ForceRemove(gid string) (g string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ForceRemove, params, &g)
	id.unlock()
	return
}

// `aria2.pause([secret, ]gid)`
// This method pauses the download denoted by gid (string).
// The status of paused download becomes paused.
// If the download was active, the download is placed in the front of waiting queue.
// While the status is paused, the download is not started.
// To change status to waiting, use the aria2.unpause() method.
// This method returns GID of paused download.
func (id *client) Pause(gid string) (g string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2Pause, params, &g)
	id.unlock()
	return
}

// `aria2.pauseAll([secret])`
// This method is equal to calling aria2.pause() for every active/waiting download.
// This methods returns OK.
func (id *client) PauseAll() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2PauseAll, params, &ok)
	id.unlock()
	return
}

// `aria2.forcePause([secret, ]gid)`
// This method pauses the download denoted by gid.
// This method behaves just like aria2.pause() except that this method pauses downloads without performing any actions which take time, such as contacting BitTorrent trackers to unregister the download first.
func (id *client) ForcePause(gid string) (g string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ForcePause, params, &g)
	id.unlock()
	return
}

// `aria2.forcePauseAll([secret])`
// This method is equal to calling aria2.forcePause() for every active/waiting download.
// This methods returns OK.
func (id *client) ForcePauseAll() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ForcePauseAll, params, &ok)
	id.unlock()
	return
}

// `aria2.unpause([secret, ]gid)`
// This method changes the status of the download denoted by gid (string) from paused to waiting, making the download eligible to be restarted.
// This method returns the GID of the unpaused download.
func (id *client) Unpause(gid string) (g string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2Unpause, params, &g)
	id.unlock()
	return
}

// `aria2.unpauseAll([secret])`
// This method is equal to calling aria2.unpause() for every active/waiting download.
// This methods returns OK.
func (id *client) UnpauseAll() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2UnpauseAll, params, &ok)
	id.unlock()
	return
}

// `aria2.tellStatus([secret, ]gid[, keys])`
// This method returns the progress of the download denoted by gid (string).
// keys is an array of strings.
// If specified, the response contains only keys in the keys array.
// If keys is empty or omitted, the response contains all keys.
// This is useful when you just want specific keys and avoid unnecessary transfers.
// For example, aria2.tellStatus("2089b05ecca3d829", ["gid", "status"]) returns the gid and status keys only.
// The response is a struct and contains following keys. Values are strings.
// https://aria2.github.io/manual/en/html/aria2c.html#aria2.tellStatus
func (id *client) TellStatus(gid string, keys ...string) (info StatusInfo, err error) {
	params := make([]interface{}, 1, 2)
	params[0] = gid
	if keys != nil {
		params = append(params, keys)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2TellStatus, params, &info)
	id.unlock()
	return
}

// `aria2.getUris([secret, ]gid)`
// This method returns the URIs used in the download denoted by gid (string).
// The response is an array of structs and it contains following keys. Values are string.
// 	uri        URI
// 	status    'used' if the URI is in use. 'waiting' if the URI is still waiting in the queue.
func (id *client) GetURIs(gid string) (infos []URIInfo, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetURIs, params, &infos)
	id.unlock()
	return
}

// `aria2.getFiles([secret, ]gid)`
// This method returns the file list of the download denoted by gid (string).
// The response is an array of structs which contain following keys. Values are strings.
// https://aria2.github.io/manual/en/html/aria2c.html#aria2.getFiles
func (id *client) GetFiles(gid string) (infos []FileInfo, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetFiles, params, &infos)
	id.unlock()
	return
}

// `aria2.getPeers([secret, ]gid)`
// This method returns a list peers of the download denoted by gid (string).
// This method is for BitTorrent only.
// The response is an array of structs and contains the following keys. Values are strings.
// https://aria2.github.io/manual/en/html/aria2c.html#aria2.getPeers
func (id *client) GetPeers(gid string) (infos []PeerInfo, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetPeers, params, &infos)
	id.unlock()
	return
}

// `aria2.getServers([secret, ]gid)`
// This method returns currently connected HTTP(S)/FTP/SFTP servers of the download denoted by gid (string).
// The response is an array of structs and contains the following keys. Values are strings.
// https://aria2.github.io/manual/en/html/aria2c.html#aria2.getServers
func (id *client) GetServers(gid string) (infos []ServerInfo, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetServers, params, &infos)
	id.unlock()
	return
}

// `aria2.tellActive([secret][, keys])`
// This method returns a list of active downloads.
// The response is an array of the same structs as returned by the aria2.tellStatus() method.
// For the keys parameter, please refer to the aria2.tellStatus() method.
func (id *client) TellActive(keys ...string) (infos []StatusInfo, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	// params = append(params, keys...)
	id.lock()
	err = Call(id.uri, aria2TellActive, params, &infos)
	id.unlock()
	return
}

// `aria2.tellWaiting([secret, ]offset, num[, keys])`
// This method returns a list of waiting downloads, including paused ones.
// offset is an integer and specifies the offset from the download waiting at the front.
// num is an integer and specifies the max. number of downloads to be returned.
// For the keys parameter, please refer to the aria2.tellStatus() method.
// If offset is a positive integer, this method returns downloads in the range of [offset, offset + num).
// offset can be a negative integer. offset == -1 points last download in the waiting queue and offset == -2 points the download before the last download, and so on.
// Downloads in the response are in reversed order then.
// For example, imagine three downloads "A","B" and "C" are waiting in this order.
// aria2.tellWaiting(0, 1) returns ["A"].
// aria2.tellWaiting(1, 2) returns ["B", "C"].
// aria2.tellWaiting(-1, 2) returns ["C", "B"].
// The response is an array of the same structs as returned by aria2.tellStatus() method.
func (id *client) TellWaiting(offset, num int, keys ...string) (infos []StatusInfo, err error) {
	params := make([]interface{}, 2, 3)
	params[0] = offset
	params[1] = num
	if keys != nil {
		params = append(params, keys)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2TellWaiting, params, &infos)
	id.unlock()
	return
}

// `aria2.tellStopped([secret, ]offset, num[, keys])`
// This method returns a list of stopped downloads.
// offset is an integer and specifies the offset from the least recently stopped download.
// num is an integer and specifies the max. number of downloads to be returned.
// For the keys parameter, please refer to the aria2.tellStatus() method.
// offset and num have the same semantics as described in the aria2.tellWaiting() method.
// The response is an array of the same structs as returned by the aria2.tellStatus() method.
func (id *client) TellStopped(offset, num int, keys ...string) (infos []StatusInfo, err error) {
	params := make([]interface{}, 2, 3)
	params[0] = offset
	params[1] = num
	if keys != nil {
		params = append(params, keys)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2TellStopped, params, &infos)
	id.unlock()
	return
}

// `aria2.changePosition([secret, ]gid, pos, how)`
// This method changes the position of the download denoted by gid in the queue.
// pos is an integer. how is a string.
// If how is POS_SET, it moves the download to a position relative to the beginning of the queue.
// If how is POS_CUR, it moves the download to a position relative to the current position.
// If how is POS_END, it moves the download to a position relative to the end of the queue.
// If the destination position is less than 0 or beyond the end of the queue, it moves the download to the beginning or the end of the queue respectively.
// The response is an integer denoting the resulting position.
// For example, if GID#2089b05ecca3d829 is currently in position 3, aria2.changePosition('2089b05ecca3d829', -1, 'POS_CUR') will change its position to 2. Additionally aria2.changePosition('2089b05ecca3d829', 0, 'POS_SET') will change its position to 0 (the beginning of the queue).
func (id *client) ChangePosition(gid string, pos int, how string) (p int, err error) {
	params := make([]interface{}, 3)
	params[0] = gid
	params[1] = pos
	params[2] = how
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ChangePosition, params, &p)
	id.unlock()
	return
}

// `aria2.changeUri([secret, ]gid, fileIndex, delUris, addUris[, position])`
// This method removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// delUris and addUris are lists of strings.
// A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// position is used to specify where URIs are inserted in the existing waiting URI list. position is 0-based.
// When position is omitted, URIs are appended to the back of the list.
// This method first executes the removal and then the addition.
// position is the position after URIs are removed, not the position when this method is called.
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris.
// In other words, if there are three URIs http://example.org/aria2 and you want remove them all, you have to specify (at least) 3 http://example.org/aria2 in delUris.
// This method returns a list which contains two integers.
// The first integer is the number of URIs deleted.
// The second integer is the number of URIs added.
func (id *client) ChangeURI(gid string, fileindex int, delUris []string, addUris []string, position ...int) (p []int, err error) {
	params := make([]interface{}, 4, 5)
	params[0] = gid
	params[1] = fileindex
	params[2] = delUris
	params[3] = addUris
	if position != nil {
		params = append(params, position[0])
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ChangeURI, params, &p)
	id.unlock()
	return
}

// `aria2.getOption([secret, ]gid)`
// This method returns options of the download denoted by gid.
// The response is a struct where keys are the names of options.
// The values are strings.
// Note that this method does not return options which have no default value and have not been set on the command-line, in configuration files or RPC methods.
func (id *client) GetOption(gid string) (m Option, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetOption, params, &m)
	id.unlock()
	return
}

// `aria2.changeOption([secret, ]gid, options)`
// This method changes options of the download denoted by gid (string) dynamically. options is a struct.
// The following options are available for active downloads:
// 	bt-max-peers
// 	bt-request-peer-speed-limit
// 	bt-remove-unselected-file
// 	force-save
// 	max-download-limit
// 	max-upload-limit
// For waiting or paused downloads, in addition to the above options, options listed in Input File subsection are available, except for following options: dry-run, metalink-base-uri, parameterized-uri, pause, piece-length and rpc-save-upload-metadata option.
// This method returns OK for success.
func (id *client) ChangeOption(gid string, option Option) (ok string, err error) {
	params := make([]interface{}, 1, 2)
	params[0] = gid
	if option != nil {
		params = append(params, option)
	}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ChangeOption, params, &ok)
	id.unlock()
	return
}

// `aria2.getGlobalOption([secret])`
// This method returns the global options.
// The response is a struct.
// Its keys are the names of options.
// Values are strings.
// Note that this method does not return options which have no default value and have not been set on the command-line, in configuration files or RPC methods. Because global options are used as a template for the options of newly added downloads, the response contains keys returned by the aria2.getOption() method.
func (id *client) GetGlobalOption() (m Option, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetGlobalOption, params, &m)
	id.unlock()
	return
}

// `aria2.changeGlobalOption([secret, ]options)`
// This method changes global options dynamically.
// options is a struct.
// The following options are available:
// 	bt-max-open-files
// 	download-result
// 	log
// 	log-level
// 	max-concurrent-downloads
// 	max-download-result
// 	max-overall-download-limit
// 	max-overall-upload-limit
// 	save-cookies
// 	save-session
// 	server-stat-of
// In addition, options listed in the Input File subsection are available, except for following options: checksum, index-out, out, pause and select-file.
// With the log option, you can dynamically start logging or change log file.
// To stop logging, specify an empty string("") as the parameter value.
// Note that log file is always opened in append mode.
// This method returns OK for success.
func (id *client) ChangeGlobalOption(options Option) (ok string, err error) {
	params := []interface{}{options}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ChangeGlobalOption, params, &ok)
	id.unlock()
	return
}

// `aria2.getGlobalStat([secret])`
// This method returns global statistics such as the overall download and upload speeds.
// The response is a struct and contains the following keys. Values are strings.
// 	downloadSpeed      Overall download speed (byte/sec).
// 	uploadSpeed        Overall upload speed(byte/sec).
// 	numActive          The number of active downloads.
// 	numWaiting         The number of waiting downloads.
// 	numStopped         The number of stopped downloads in the current session.
//                     This value is capped by the --max-download-result option.
// 	numStoppedTotal    The number of stopped downloads in the current session and not capped by the --max-download-result option.
func (id *client) GetGlobalStat() (info GlobalStatInfo, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetGlobalStat, params, &info)
	id.unlock()
	return
}

// `aria2.purgeDownloadResult([secret])`
// This method purges completed/error/removed downloads to free memory.
// This method returns OK.
func (id *client) PurgeDownloadResult() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2PurgeDownloadResult, params, &ok)
	id.unlock()
	return
}

// `aria2.removeDownloadResult([secret, ]gid)`
// This method removes a completed/error/removed download denoted by gid from memory.
// This method returns OK for success.
func (id *client) RemoveDownloadResult(gid string) (ok string, err error) {
	params := []interface{}{gid}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2RemoveDownloadResult, params, &ok)
	id.unlock()
	return
}

// `aria2.getVersion([secret])`
// This method returns the version of aria2 and the list of enabled features.
// The response is a struct and contains following keys.
// 	version            Version number of aria2 as a string.
// 	enabledFeatures    List of enabled features. Each feature is given as a string.
func (id *client) GetVersion() (info VersionInfo, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetVersion, params, &info)
	id.unlock()
	return
}

// `aria2.getSessionInfo([secret])`
// This method returns session information.
// The response is a struct and contains following key.
// 	sessionId    Session ID, which is generated each time when aria2 is invoked.
func (id *client) GetSessionInfo() (info SessionInfo, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2GetSessionInfo, params, &info)
	id.unlock()
	return
}

// `aria2.shutdown([secret])`
// This method shutdowns aria2.
// This method returns OK.
func (id *client) Shutdown() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2Shutdown, params, &ok)
	id.unlock()
	return
}

// `aria2.forceShutdown([secret])`
// This method shuts down aria2().
// This method behaves like :func:'aria2.shutdown` without performing any actions which take time, such as contacting BitTorrent trackers to unregister downloads first.
// This method returns OK.
func (id *client) ForceShutdown() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2ForceShutdown, params, &ok)
	id.unlock()
	return
}

// `aria2.saveSession([secret])`
// This method saves the current session to a file specified by the --save-session option.
// This method returns OK if it succeeds.
func (id *client) SaveSession() (ok string, err error) {
	params := []interface{}{}
	if id.token != "" {
		params = append(params, "token:"+id.token)
	}
	id.lock()
	err = Call(id.uri, aria2SaveSession, params, &ok)
	id.unlock()
	return
}

// `system.multicall(methods)`
// This methods encapsulates multiple method calls in a single request.
// methods is an array of structs.
// The structs contain two keys: methodName and params.
// methodName is the method name to call and params is array containing parameters to the method call.
// This method returns an array of responses.
// The elements will be either a one-item array containing the return value of the method call or a struct of fault element if an encapsulated method call fails.
func (id *client) Multicall(methods []Method) (r []interface{}, err error) {
	if len(methods) == 0 {
		err = errInvalidParameter
		return
	}
	id.lock()
	err = Call(id.uri, aria2Multicall, methods, &r)
	id.unlock()
	return
}

// `system.listMethods()`
// This method returns the all available RPC methods in an array of string.
// Unlike other methods, this method does not require secret token.
// This is safe because this method jsut returns the available method names.
func (id *client) ListMethods() (methods []string, err error) {
	id.lock()
	err = Call(id.uri, aria2ListMethods, []interface{}{}, &methods)
	id.unlock()
	return
}
