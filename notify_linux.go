package notify

import (
	"os/exec"
	"path/filepath"
)

// Some vars
var (
	// The sound backend to use
	SoundBackEnd string

	// Array of possible sound backends to use
	SoundBackEnds []string = []string{

		"ogg123",
		"paplay",
		//"aplay",
	}
)

func init() {

	// Look for a suitable sound backend
	for _, v := range SoundBackEnds {

		path, _ := exec.LookPath(v)
		if len(path) > 1 {

			SoundBackEnd = v
		}
	}
}

func Send(title, body string) (err error) {

	m := &Message{

		Title: title,
		Body:  body,
	}
	return m.Send()
}

func (m *Message) Send() (err error) {

	var args []string

	if (len(m.Title) == 0) && (len(m.Body) == 0) {

		return ErrTitleMsg
	}

	if len(m.Title) > 0 {

		args = append(args, m.Title)
	}

	if len(m.Body) > 0 {

		args = append(args, m.Body)
	}

	if len(m.Icon) > 0 {

		args = append(args, "--icon="+filepath.Clean(m.Icon))
	}

	if len(m.Urgency) > 0 {

		args = append(args, "--urgency="+m.Urgency)
	}

	if m.ExpireTime > 0 {

		args = append(args, "--expire-time="+string(m.ExpireTime))
	}

	if len(m.Category) > 0 {

		args = append(args, "--category="+m.Category)
	}

	if len(m.Hint) > 0 {

		args = append(args, "--hint="+m.Hint)
	}

	out, err := exec.Command("notify-send", args...).CombinedOutput()
	if err != nil {

		return &Error{Return: string(out), Err: err}
	}

	if (len(m.Sound) > 0) || (len(m.SoundPipe) > 0) {

		return m.PlaySound()
	}

	return
}

func (m *Message) PlaySound() (err error) {

	if len(SoundBackEnd) == 0 {

		return ErrNoSoundBackEnd
	}

	if len(m.Sound) > 0 {

		out, err := exec.Command(SoundBackEnd, m.Sound).CombinedOutput()
		if err != nil {

			return &Error{Return: string(out), Err: err}
		}

	} else if len(m.SoundPipe) > 0 {

		cmd := exec.Command(SoundBackEnd, "-")

		stdin, err := cmd.StdinPipe()
		if err != nil {

			return err
		}

		if _, err = stdin.Write(m.SoundPipe); err != nil {

			return err
		}

		if err = stdin.Close(); err != nil {

			return err
		}

		if err := cmd.Run(); err != nil {

			return err
		}
	}

	return
}
