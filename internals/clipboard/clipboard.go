package clipboard

import (
  "os/exec"
  "io"
  "bytes"
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
	cmd := exec.Command("xclip", "-selection", "clipboard")
  cmd.Stdin = io.NopCloser(bytes.NewReader(data))

  if err := cmd.Run(); err != nil {
    return err
  }

  return nil
}
