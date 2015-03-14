package notify

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
