package main

import (
	"log"

	"github.com/TheCreeper/go-notify"
)

func main() {

	n, err := notify.NewNotification("Test Notification", "This is a test notification")
	if err != nil {

		log.Fatal(err)
	}

	if _, err := n.Send(); err != nil {

		log.Fatal(err)
	}
}
