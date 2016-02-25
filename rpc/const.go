package rpc

const (
	addURI               = "aria2.addUri"
	addTorrent           = "aria2.addTorrent"
	addMetalink          = "aria2.addMetalink"
	remove               = "aria2.remove"
	forceRemove          = "aria2.forceRemove"
	pause                = "aria2.pause"
	pauseAll             = "aria2.pauseAll"
	forcePause           = "aria2.forcePause"
	forcePauseAll        = "aria2.forcePauseAll"
	unpause              = "aria2.unpause"
	unpauseAll           = "aria2.unpauseAll"
	tellStatus           = "aria2.tellStatus"
	getUris              = "aria2.getUris"
	getFiles             = "aria2.getFiles"
	getPeers             = "aria2.getPeers"
	getServers           = "aria2.getServers"
	tellActive           = "aria2.tellActive"
	tellWaiting          = "aria2.tellWaiting"
	tellStopped          = "aria2.tellStopped"
	changePosition       = "aria2.changePosition"
	changeURI            = "aria2.changeUri"
	getOption            = "aria2.getOption"
	changeOption         = "aria2.changeOption"
	getGlobalOption      = "aria2.getGlobalOption"
	changeGlobalOption   = "aria2.changeGlobalOption"
	getGlobalStat        = "aria2.getGlobalStat"
	purgeDowloadResult   = "aria2.purgeDowloadResult"
	removeDownloadResult = "aria2.removeDownloadResult"
	getVersion           = "aria2.getVersion"
	getSessionInfo       = "aria2.getSessionInfo"
	shutdown             = "aria2.shutdown"
	forceShutdown        = "aria2.forceShutdown"
	multicall            = "system.multicall"
)