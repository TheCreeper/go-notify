package notify

import (
	"errors"

	"github.com/godbus/dbus"
)

// Returns the capabilities of the notification server as a string array
func GetCapabilities() (c []string, err error) {

	connection, err := dbus.SessionBus()
	if err != nil {

		return
	}

	call := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications").Call("org.freedesktop.Notifications.GetCapabilities", 0)

	if call.Err != nil {

		return nil, call.Err
	}

	if err := call.Store(&c); err != nil {

		return nil, err
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

	call := connection.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications").Call("org.freedesktop.Notifications.GetCapabilities", 0, id)

	if call.Err != nil {

		return call.Err
	}

	return
}
