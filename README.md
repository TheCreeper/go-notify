go-notify
=====================

[![go-notify](https://godoc.org/github.com/TheCreeper/go-notify?status.png)](http://godoc.org/github.com/TheCreeper/go-notify)

This is a golang library that provides an implementation of the [freedesktop notification spec](https://developer.gnome.org/notification-spec/) using the DBUS api.

## Example

```
package main

import (
	"log"

	"github.com/TheCreeper/go-notify"
)

func main() {

	ntf, err := notify.NewNotification("Test Notification", "This is a test notification")
	if err != nil {

		log.Fatal(err)
	}

	if _, err := ntf.Send(); err != nil {

		log.Fatal(err)
	}
}
```