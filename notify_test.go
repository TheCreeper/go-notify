package notify

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

func TestGetCapabilities(t *testing.T) {
	c, err := GetCapabilities()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Support Action Icons: %v\n", c.ActionIcons)
	t.Logf("Support Actions: %v\n", c.Actions)
	t.Logf("Support Body: %v\n", c.Body)
	t.Logf("Support Body Hyperlinks: %v\n", c.BodyHyperlinks)
	t.Logf("Support Body Images: %v\n", c.BodyImages)
	t.Logf("Support Body Markup: %v\n", c.BodyMarkup)
	t.Logf("Support Icon Multi: %v\n", c.IconMulti)
	t.Logf("Support Icon Static: %v\n", c.IconStatic)
	t.Logf("Support Persistence: %v\n", c.Persistence)
	t.Logf("Support Sound: %v\n", c.Sound)
}

func TestGetServerInformation(t *testing.T) {
	info, err := GetServerInformation()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Server Name: %s\n", info.Name)
	t.Logf("Server Spec Version: %s\n", info.SpecVersion)
	t.Logf("Server Vendor: %s\n", info.Vendor)
	t.Logf("Sserver Version: %s\n", info.Version)
}

func TestNewNotification(t *testing.T) {
	ntf := NewNotification("Notification Test", "Just a test")
	if _, err := ntf.Show(); err != nil {
		t.Fatal(err)
	}
}

func TestCloseNotification(t *testing.T) {
	ntf := NewNotification("Notification Test", "Just a test")
	id, err := ntf.Show()
	if err != nil {
		t.Fatal(err)
	}

	if err = CloseNotification(id); err != nil {
		t.Fatal(err)
	}
}

func TestUrgencyNotification(t *testing.T) {
	ntfLow := NewNotification("Urgency Test", "Testing notification urgency low")
	ntfLow.Hints = make(map[string]interface{})

	ntfLow.Hints[HintUrgency] = UrgencyLow
	_, err := ntfLow.Show()
	if err != nil {
		t.Fatal(err)
	}

	ntfCritical := NewNotification("Urgency Test", "Testing notification urgency critical")
	ntfCritical.Hints = make(map[string]interface{})

	ntfCritical.Hints[HintUrgency] = UrgencyCritical
	_, err = ntfCritical.Show()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignalNotify(t *testing.T) {
	msg := "Just a test\n\n<b>you can click things!</b>\n\n(use the -short test switch to skip this)"
	ntf := NewNotification("Notification Test", msg)
	if testing.Short() {
		ntf.Timeout = 500
		log.Print("Using short timeout because of -short")
	} else {
		ntf.Timeout = 5000
		log.Printf("Using %vms timeout, click the notification or use -short to go faster", ntf.Timeout)
	}
	ntf.Actions = []string{"my-something", "Something", "default", "Default"}
	id, err := ntf.Show()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	ch := make(chan Signal, 2)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for sig := range ch {
			log.Printf("Signal: %+v", sig)
			if sig.CloseReason == NotClosed {
				CloseNotification(id)
			} else {
				cancel()
				break
			}
		}
	}()

	err = SignalNotify(ctx, ch)
	if err != nil && err != context.Canceled {
		t.Fatal(err)
	}

	close(ch)
	wg.Wait()
}
