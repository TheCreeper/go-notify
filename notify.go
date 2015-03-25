package notify

import (
	"errors"

	"github.com/godbus/dbus"
)

// Notification expire times
const (
	ExpiresDefault = -1
	ExpiresNever   = 0
)

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
	UrgencyLow      = 0
	UrgencyNormal   = 1
	UrgencyCritical = 2
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

type Capabilities struct {
	ActionIcons    bool
	Actions        bool
	Body           bool
	BodyHyperlinks bool
	BodyImages     bool
	BodyMarkup     bool
	IconMulti      bool
	IconStatic     bool
	Persistence    bool
	Sound          bool
}

// Returns the capabilities of the notification server
func GetCapabilities() (c *Capabilities, err error) {

	connection, err := dbus.SessionBus()
	if err != nil {

		return
	}
	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.GetCapabilities", 0)
	if call.Err != nil {

		return nil, call.Err
	}

	c = &Capabilities{}
	s := []string{}

	if err := call.Store(&s); err != nil {

		return nil, err
	}

	for _, v := range s {

		if v == ActionIcons {

			c.ActionIcons = true
			continue
		}

		if v == Actions {

			c.Actions = true
			continue
		}

		if v == Body {

			c.Body = true
			continue
		}

		if v == BodyHyperlinks {

			c.BodyHyperlinks = true
			continue
		}

		if v == BodyImages {

			c.BodyImages = true
			continue
		}

		if v == BodyMarkup {

			c.BodyMarkup = true
			continue
		}

		if v == IconMulti {

			c.IconMulti = true
			continue
		}

		if v == IconStatic {

			c.IconStatic = true
			continue
		}

		if v == Persistence {

			c.Persistence = true
			continue
		}

		if v == Sound {

			c.Sound = true
			continue
		}
	}

	return
}

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

// Gets information about the notification server
func GetServerInformation() (i *ServerInformation, err error) {

	connection, err := dbus.SessionBus()
	if err != nil {

		return
	}
	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.GetServerInformation", 0)
	if call.Err != nil {

		return nil, call.Err
	}

	i = &ServerInformation{}
	if err := call.Store(&i.Name, &i.Vendor, &i.Version, &i.SpecVersion); err != nil {

		return nil, err
	}

	return
}

type Notification struct {

	// The optional name of the application sending the notification. Can be blank.
	AppName string

	// The optional notification ID that this notification replaces.
	ReplacesID uint32

	// The optional program icon of the calling application.
	AppIcon string

	// The summary text briefly describing the notification.
	Summary string

	// The optional detailed body text.
	Body string

	// The actions send a request message back to the notification client when invoked.
	Actions []string

	// Optional hints that can be passed to the server from the client program.
	Hints map[string]string

	// The timeout time in milliseconds since the display of the notification at which the notification should automatically close.
	Timeout int32
}

// Create a new notification object with some basic information
func NewNotification(summary, body string) (n *Notification, err error) {

	if len(summary) == 0 {

		return nil, errors.New("The Notification must contain a summary")
	}

	n = &Notification{

		Summary: summary,
		Body:    body,
		Timeout: ExpiresDefault,
	}
	return
}

// Send the notification
func (n *Notification) Show() (id uint32, err error) {

	hints := map[string]dbus.Variant{}
	if len(n.Hints) != 0 {

		for k, v := range n.Hints {

			hints[k] = dbus.MakeVariant(v)
		}
	}

	connection, err := dbus.SessionBus()
	if err != nil {

		return
	}
	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		n.AppName,
		n.ReplacesID,
		n.AppIcon,
		n.Summary,
		n.Body,
		n.Actions,
		hints,
		n.Timeout)
	if call.Err != nil {

		return 0, call.Err
	}

	if err := call.Store(&id); err != nil {

		return 0, err
	}

	return
}

// Closes the notification if it exists using its id
func CloseNotification(id uint32) (err error) {

	connection, err := dbus.SessionBus()
	if err != nil {

		return
	}
	obj := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.CloseNotification", 0, id)
	if call.Err != nil {

		return call.Err
	}

	return
}
