package notify

import (
	"fmt"
	"testing"
)

func TestGetCapabilities(t *testing.T) {

	c, err := GetCapabilities()
	if err != nil {

		t.Error(err)
	}

	fmt.Printf("Support Action Icons: %v\n", c.ActionIcons)
	fmt.Printf("Support Actions: %v\n", c.Actions)
	fmt.Printf("Support Body: %v\n", c.Body)
	fmt.Printf("Support Body Hyperlinks: %v\n", c.BodyHyperlinks)
	fmt.Printf("Support Body Images: %v\n", c.BodyImages)
	fmt.Printf("Support Body Markup: %v\n", c.BodyMarkup)
	fmt.Printf("Support Icon Multi: %v\n", c.IconMulti)
	fmt.Printf("Support Icon Static: %v\n", c.IconStatic)
	fmt.Printf("Support Persistence: %v\n", c.Persistence)
	fmt.Printf("Support Sound: %v\n", c.Sound)
}

func TestGetServerInformation(t *testing.T) {

	info, err := GetServerInformation()
	if err != nil {

		t.Error(err)
	}

	fmt.Printf("Server Name: %s\n", info.Name)
	fmt.Printf("Server Spec Version: %s\n", info.SpecVersion)
	fmt.Printf("Server Vendor: %s\n", info.Vendor)
	fmt.Printf("Sserver Version: %s\n", info.Version)
}

func TestNewNotification(t *testing.T) {

	ntf, err := NewNotification("Test Notification", "This is a test notification")
	if err != nil {

		t.Error(err)
	}

	id, err := ntf.Show()
	if err != nil {

		t.Error(err)
	}

	if err := CloseNotification(id); err != nil {

		t.Error(err)
	}
}
