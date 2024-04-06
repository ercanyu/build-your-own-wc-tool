package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

const DataDirectory = "/data/"

func OpenFile(fileName string) (*os.File, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, err
	}
	fullFilePath := projectRoot + DataDirectory + fileName
	file, err := os.Open(fullFilePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func findProjectRoot() (string, error) {
	markerFileName := "go.mod"
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, markerFileName)); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("marker file '%s' not found in the directory tree", markerFileName)
		}

		currentDir = parentDir
	}
}
