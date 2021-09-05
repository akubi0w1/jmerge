package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile reads specified file.
func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// WriteFile writes file.
// If directories of files contained in path dones not exist, create them.
func WriteFile(path, fileName string, content []byte) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	fullPath := filepath.Join(path, fileName)
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := ioutil.WriteFile(fullPath, content, 0755); err != nil {
		return err
	}
	return nil
}
