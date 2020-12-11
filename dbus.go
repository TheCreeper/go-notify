package notify

import (
	"context"
	"github.com/godbus/dbus/v5"
	"image"
)

// Notification object paths and interfaces.
const (
	DbusObjectPath               = "/org/freedesktop/Notifications"
	DbusInterfacePath            = "org.freedesktop.Notifications"
	SignalNotificationClosed     = "org.freedesktop.Notifications.NotificationClosed"
	SignalActionInvoked          = "org.freedesktop.Notifications.ActionInvoked"
	CallGetCapabilities          = "org.freedesktop.Notifications.GetCapabilities"
	CallCloseNotification        = "org.freedesktop.Notifications.CloseNotification"
	CallNotify                   = "org.freedesktop.Notifications.Notify"
	CallGetServerInformation     = "org.freedesktop.Notifications.GetServerInformation"
	DbusMemberActionInvoked      = "ActionInvoked"
	DbusMemberNotificationClosed = "NotificationClosed"
)

// Notification expire timeout.
const (
	ExpiresDefault = -1
	ExpiresNever   = 0
)

// Notification Categories
const (
	ClassDevice              = "device"
	ClassDeviceAdded         = "device.added"
	ClassDeviceError         = "device.error"
	ClassDeviceRemoved       = "device.removed"
	ClassEmail               = "email"
	ClassEmailArrived        = "email.arrived"
	ClassEmailBounced        = "email.bounced"
	ClassIm                  = "im"
	ClassImError             = "im.error"
	ClassImReceived          = "im.received"
	ClassNetwork             = "network"
	ClassNetworkConnected    = "network.connected"
	ClassNetworkDisconnected = "network.disconnected"
	ClassNetworkError        = "network.error"
	ClassPresence            = "presence"
	ClassPresenceOffline     = "presence.offline"
	ClassPresenceOnline      = "presence.online"
	ClassTransfer            = "transfer"
	ClassTransferComplete    = "transfer.complete"
	ClassTransferError       = "transfer.error"
)

// Urgency Levels
const (
	UrgencyLow      = byte(0)
	UrgencyNormal   = byte(1)
	UrgencyCritical = byte(2)
)

// Hints
const (
	HintActionIcons   = "action-icons"
	HintCategory      = "category"
	HintDesktopEntry  = "desktop-entry"
	HintImageData     = "image-data"
	HintImagePath     = "image-path"
	HintResident      = "resident"
	HintSoundFile     = "sound-file"
	HintSoundName     = "sound-name"
	HintSuppressSound = "suppress-sound"
	HintTransient     = "transient"
	HintX             = "x"
	HintY             = "y"
	HintUrgency       = "urgency"
)

// Capabilities is a struct containing the capabilities of the notification
// server.
type Capabilities struct {
	// Supports using icons instead of text for displaying actions.
	ActionIcons bool

	// The server will provide any specified actions to the user.
	Actions bool

	// Supports body text. Some implementations may only show the summary.
	Body bool

	// The server supports hyperlinks in the notifications.
	BodyHyperlinks bool

	// The server supports images in the notifications.
	BodyImages bool

	// Supports markup in the body text.
	BodyMarkup bool

	// The server will render an animation of all the frames in a given
	// image array.
	IconMulti bool

	// Supports display of exactly 1 frame of any given image array.
	IconStatic bool

	// The server supports persistence of notifications. Notifications will
	// be retained until they are acknowledged or removed by the user or
	// recalled by the sender.
	Persistence bool

	// The server supports sounds on notifications.
	Sound bool
}

// GetCapabilities returns the capabilities of the notification server.
func GetCapabilities() (c Capabilities, err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}

	var d = make(chan *dbus.Call, 1)
	var o = conn.Object(DbusInterfacePath, DbusObjectPath)
	var s = make([]string, 0)
	o.GoWithContext(context.Background(),
		CallGetCapabilities,
		0,
		d)
	err = (<-d).Store(&s)
	if err != nil {
		return
	}

	for _, v := range s {
		switch v {
		case "action-icons":
			c.ActionIcons = true
			break
		case "actions":
			c.Actions = true
			break
		case "body":
			c.Body = true
			break
		case "body-hyperlinks":
			c.BodyHyperlinks = true
			break
		case "body-images":
			c.BodyImages = true
			break
		case "body-markup":
			c.BodyMarkup = true
			break
		case "icon-multi":
			c.IconMulti = true
			break
		case "icon-static":
			c.IconStatic = true
			break
		case "persistence":
			c.Persistence = true
			break
		case "sound":
			c.Sound = true
			break
		}
	}
	return
}

// ServerInformation is a struct containing information about the server such
// as its name and version.
type ServerInformation struct {
	// The name of the notification server daemon
	Name string

	// The vendor of the notification server
	Vendor string

	// Version of the notification server
	Version string

	// Spec version the notification server conforms to
	SpecVersion string
}

// GetServerInformation returns information about the notification server such
// as its name and version.
func GetServerInformation() (i ServerInformation, err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}
	var d = make(chan *dbus.Call, 1)
	var o = conn.Object(DbusInterfacePath, DbusObjectPath)
	o.GoWithContext(context.Background(),
		CallGetServerInformation,
		0,
		d)
	err = (<-d).Store(&i.Name,
		&i.Vendor,
		&i.Version,
		&i.SpecVersion)
	return
}

// Notification is a struct which describes the notification to be displayed
// by the notification server.
type Notification struct {
	// The optional name of the application sending the notification.
	// Can be blank.
	AppName string

	// The optional notification ID that this notification replaces.
	ReplacesID uint32

	// The optional program icon of the calling application.
	AppIcon string

	// The summary text briefly describing the notification.
	Summary string

	// The optional detailed body text.
	Body string

	// The actions send a request message back to the notification client
	// when invoked.
	Actions []string

	// Hints are a way to provide extra data to a notification server.
	Hints map[string]interface{}

	// The timeout time in milliseconds since the display of the
	// notification at which the notification should automatically close.
	Timeout int32

	hints map[string]dbus.Variant
}

// NewNotification creates a new notification object with some basic
// information.
func NewNotification(summary, body string) Notification {
	return Notification{
		Body:    body,
		Summary: summary,
		Timeout: ExpiresDefault,
		hints:   make(map[string]dbus.Variant),
	}
}

// Show sends the information in the notification object to the server to be
// displayed.
func (x *Notification) Show() (id uint32, err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}

	// We need to convert the interface type of the map to dbus.Variant as
	// people dont want to have to import the dbus package just to make use
	// of the notification hints.
	for k, v := range x.Hints {
		x.hints[k] = dbus.MakeVariant(v)
	}

	var d = make(chan *dbus.Call, 1)
	var o = conn.Object(DbusInterfacePath, DbusObjectPath)
	o.GoWithContext(context.Background(),
		CallNotify,
		0,
		d,
		x.AppName,
		x.ReplacesID,
		x.AppIcon,
		x.Summary,
		x.Body,
		x.Actions,
		x.hints,
		x.Timeout)
	err = (<-d).Store(&id)
	return
}

// _ImageData needs to resemble the (iiibiiay) signature.
type _ImageData struct {
	/*0*/ Width int
	/*1*/ Height int
	/*2*/ RowStride int
	/*3*/ HasAlpha bool
	/*4*/ BitsPerSample int
	/*5*/ Samples int
	/*6*/ Image []byte
}

type ImageError struct{}

func (x ImageError) Error() string {
	return "Given image.Image was not of type *image.RGBA"
}

// SetImage sets the image in the notification from an image.Image
// interface which must have an underlying type of *image.RGBA.
// Only the RGBA color space is allowed as only that is supported
// by the gdk-pixbuf library.
func (x *Notification) SetImage(img image.Image) (err error) {
	if p, ok := img.(*image.RGBA); ok {
		var r = p.Bounds()
		var d = _ImageData{
			r.Max.X, // Width
			r.Max.Y, // Height
			p.Stride,
			true,
			8,
			4,
			p.Pix,
		}
		x.hints["image-data"] = dbus.MakeVariant(d)
		return
	}
	err = ImageError{}
	return
}

// CloseNotification closes the notification if it exists using its id.
func CloseNotification(id uint32) (err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}
	var d = make(chan *dbus.Call, 1)
	var o = conn.Object(DbusInterfacePath, DbusObjectPath)
	o.GoWithContext(context.Background(),
		CallCloseNotification,
		0,
		d,
		id)
	err = (<-d).Err
	return
}
