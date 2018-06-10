// +build darwin !linux !windows

package notify

import (
	"fmt"
	"os/exec"
)

const format = `display notification "%s" with title "%s"`

// Message represents a notification payload.
type Message struct {
	Title string
	Body  string
}

// Alert creates a desktop notification.
func Alert(message *Message) error {
	cmd := exec.Command("osascript")

	cmd.Args = append(cmd.Args, "-e")
	cmd.Args = append(cmd.Args, fmt.Sprintf(format, message.Body, message.Title))

	return cmd.Run()
}
