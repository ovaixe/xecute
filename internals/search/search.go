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

		if info.IsDir() {
			return nil
		}

		match := false

		switch {
		case insensitive:
			match = strings.Contains(strings.ToLower(info.Name()), strings.ToLower(fileName))
		default:
			match = info.Name() == fileName
		}

		if match {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			filePaths = append(filePaths, absPath)
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
