package notify

import (
	"errors"

	"github.com/godbus/dbus"
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

// Create a new notification object with some basic information
func NewNotification(summary, body string) (n *Notification, err error) {

	if len(summary) == 0 {

		return nil, errors.New("The Notification must contain a summary")
	}

	n = &Notification{

		Summary: summary,
		Body:    body,
		Timeout: -1,
	}
	return
}

// Send the notification
func (n *Notification) Send() (id uint32, err error) {

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

	call := obj.Call("org.freedesktop.Notifications.GetCapabilities", 0, id)
	if call.Err != nil {

		return call.Err
	}

	return
}

type ServerInformation struct {
	Name        string
	Vendor      string
	Version     string
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
	if err := call.Store(&i.Name, i.Vendor, i.Version, i.SpecVersion); err != nil {

		return nil, err
	}

	return
}
