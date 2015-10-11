# go-notify

[go-notify]: https://godoc.org/github.com/TheCreeper/go-notify
[go-notify-img]: https://godoc.org/github.com/TheCreeper/go-notify?status.svg
[Notification Specification]: https://developer.gnome.org/notification-spec/
[Icon Naming Specification]: http://standards.freedesktop.org/icon-naming-spec/

[![GoDoc]([go-notify-img])]([go-notify])

Package notify provides an implementation of the Gnome DBus
[Notification Specification].

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
ntf := notify.NewNotification("Test Notification", "Just a test")
ntf.Hints = make(map[string]interface{})
ntf.Hints[notify.HintSoundFile] = "/home/my-username/sound.oga"
if _, err := ntf.Show(); err != nil {
	return
}
```