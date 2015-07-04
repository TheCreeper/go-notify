go-notify
=====================

[![go-notify](https://godoc.org/github.com/TheCreeper/go-notify?status.png)](http://godoc.org/github.com/TheCreeper/go-notify)

Package notify provides an implementation of the [Freedesktop Notifications Specification](https://developer.gnome.org/notification-spec/) using the DBus API.

## Examples

Display Simple Notification
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")

if _, err := ntf.Show(); err != nil {
	return
}
```

Display a Notification with Icon. Consult the [Icon Naming Specification](http://standards.freedesktop.org/icon-naming-spec/icon-naming-spec-latest.html).
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
//ntf.AppIcon = "/usr/share/icons/gnome/scalable/devices/network-wireless-symbolic.svg"
ntf.AppIcon = "network-wireless"

if _, err := ntf.Show(); err != nil {
	return
}
```

Display a Notification with a High Priority
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.Hints = make(map[string]interface{})
ntf.Hints[notify.Urgency] = notify.UrgencyCritical

if _, err := ntf.Show(); err != nil {
	return
}
```

Play a Sound with the Notification
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.Hints = make(map[string]interface{})
ntf.Hints[notify.HintSoundFile] = "/usr/share/sounds/freedesktop/stereo/dialog-information.oga"

if _, err := ntf.Show(); err != nil {
	return
}
```

Close a Notification
```Go
if err = notify.CloseNotification(notification_id); err != nil {
	return
}
```