package main

// Categories
const (
	Device              = "device"
	DeviceAdded         = "device.added"
	DeviceError         = "device.error"
	DeviceRemoved       = "device.removed"
	Email               = "email"
	EmailArrived        = "email.arrived"
	EmailBounced        = "email.bounced"
	Im                  = "im"
	ImError             = "im.error"
	ImReceived          = "im.received"
	Network             = "network"
	NetworkConnected    = "network.connected"
	NetworkDisconnected = "network.disconnected"
	NetworkError        = "network.error"
	Presence            = "presence"
	PresenceOffline     = "presence.offline"
	PresenceOnline      = "presence.online"
	Transfer            = "transfer"
	TransferComplete    = "transfer.complete"
	TransferError       = "transfer.error"
)

// Urgency Levels
const (
	Low      = 0
	Normal   = 1
	Critical = 2
)

// Capabilities
const (
	ActionIcons    = "action-icons"
	Actions        = "actions"
	Body           = "body"
	BodyHyperlinks = "body-hyperlinks"
	BodyImages     = "body-images"
	BodyMarkup     = "body-markup"
	IconMulti      = "icon-multi"
	IconStatic     = "icon-static"
	Persistence    = "persistence"
	Sound          = "sound"
)
