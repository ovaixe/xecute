package clipboard

import (
	"errors"
	"os/exec"

	"github.com/ovaixe/xecute/internals/platform"
)

func Read() (string, error) {
	cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func Write(data []byte) error {
	var cmd *exec.Cmd

	// Check if WAYLAND_DISPLAY is set
	if platform.IsWayland() {
		cmd = exec.Command("wl-copy")
	} else if platform.IsX11() {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	} else {
		return errors.New("no clipboard tool available: not running on wayland or X11")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = stdin.Write(data)
	if err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
