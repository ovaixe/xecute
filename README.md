# xecute

A simple CLI tool for searching files and copying file to the clipboard.

## Dependencies

- For X11 systems:

  - Requires `xclip`

- For Wayland systems:
  - Requires `wl-clipboard`

**Note:** Make sure the required clipboard utility is installed and available in your `PATH`

## Usage

Search for a file:
- `xecute s -r=/home filename`

Copy file contents to clipboard
- `xecute c -d=/home filename`

### ✨ Bonus tip

If you want to be super helpful, you can also add more features
