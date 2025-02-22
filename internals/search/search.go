package search

import (
  "errors"
  "io/fs"
  "path/filepath"
  "strings"
)

var ErrFileNotFound = errors.New("FILE NOT FOUND")

func SearchFile(dir, fileName string, insensitive bool) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		switch {
		case insensitive:
			if !info.IsDir() && strings.Contains(strings.ToLower(info.Name()), strings.ToLower(fileName)) {
				filePaths = append(filePaths, path)
			}
		default:
			if !info.IsDir() && info.Name() == fileName {
				filePaths = append(filePaths, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, ErrFileNotFound
	}

  return filePaths, nil
}
