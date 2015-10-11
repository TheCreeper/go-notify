# go-notify

[go-notify]: http://godoc.org/github.com/TheCreeper/go-notify
[go-notify-img]: https://godoc.org/github.com/TheCreeper/go-notify?status.png
[Notification Specification]: https://developer.gnome.org/notification-spec/
[Icon Naming Specification]: http://standards.freedesktop.org/icon-naming-spec/

[![go-notify-img][]]([go-notify])

Package notify provides an implementation of the [Notification Specification]
using the DBus API.

## Examples

Display a Simple Notification
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
if _, err := ntf.Show(); err != nil {
	return
}
```

Display a Notification with an Icon. Consult the [Icon Naming Specification].
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.AppIcon = "network-wireless"
if _, err := ntf.Show(); err != nil {
	return
}
```

Display a Notification that never Expires
```Go
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.Timeout = notify.ExpiresNever
if _, err := ntf.Show(); err != nil {
	return
}
```

Play a Sound with the Notification
```Go
snd := "/usr/share/sounds/freedesktop/stereo/dialog-information.oga"
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.Hints = make(map[string]interface{})
ntf.Hints[notify.HintSoundFile] = snd
if _, err := ntf.Show(); err != nil {
	return
}
```