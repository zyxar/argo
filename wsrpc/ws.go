package wsrpc

// notification
/*
aria2.onDownloadStart(event)
This notification will be sent if a download is started.
The event is of type struct and it contains following keys.
The value type is string.

gid
GID of the download.
aria2.onDownloadPause(event)
This notification will be sent if a download is paused.
The event is the same struct of the event argument of aria2.onDownloadStart() method.

aria2.onDownloadStop(event)
This notification will be sent if a download is stopped by the user.
The event is the same struct of the event argument of aria2.onDownloadStart() method.

aria2.onDownloadComplete(event)
This notification will be sent if a download is completed.
In BitTorrent downloads, this notification is sent when the download is completed and seeding is over.
The event is the same struct of the event argument of aria2.onDownloadStart() method.

aria2.onDownloadError(event)
This notification will be sent if a download is stopped due to error.
The event is the same struct of the event argument of aria2.onDownloadStart() method.

aria2.onBtDownloadComplete(event)
This notification will be sent if a download is completed in BitTorrent (but seeding may not be over).
The event is the same struct of the event argument of aria2.onDownloadStart() method.
*/
