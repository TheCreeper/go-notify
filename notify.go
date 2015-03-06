package notify

import (
	"errors"
	"strings"
)

type Message struct {
	Title      string
	Body       string
	Icon       string
	Urgency    string
	ExpireTime int
	Category   string
	Hint       string
	Sound      string
	SoundPipe  []byte
}

const (
	LowPriority      = "low"
	NormalPriority   = "normal"
	CriticalPriority = "critical"
)

// Errors
var (
	ErrTitleMsg       = errors.New("A title or message must be specified.")
	ErrNoSoundBackEnd = errors.New("No sound backend could be found.")
)

type Error struct {
	Return string
	Err    error
}

func (e *Error) Error() string {

	// Usually return will have a newline character
	return "Notify: " + " " + strings.TrimSpace(e.Return) + ": " + e.Err.Error()
}
