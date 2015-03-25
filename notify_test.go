package notify

import (
	"log"
	"testing"
)

func TestNotificationNew(t *testing.T) {

	ntf, err := NotificationNew("Test Notification", "This is a test notification")
	if err != nil {

		log.Fatal(err)
	}

	if _, err := ntf.Show(); err != nil {

		log.Fatal(err)
	}
}
