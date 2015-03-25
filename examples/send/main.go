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
