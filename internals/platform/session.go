package platform

import (
	"os"
)

// IsWayland returns true if running on a wayland session
func IsWayland() bool {
	return os.Getenv("WAYLAND_DISPLAY") != ""
}

// IsX11 returns true if running on an x11 session
func IsX11() bool {
	return os.Getenv("DISPLAY") != ""
}
