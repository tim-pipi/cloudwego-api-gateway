package dir

import (
	"io"
	"log"
	"os"
)

var _ = copyFile

// Copies the file from source to destination
func copyFile(srcPath string, destPath string) error {
	src, err := os.Open(srcPath)

	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}

// Creates a directory if it doesn't exist
func CreateDir(path string) error {
	e, err := exists(path)
	if err != nil {
		return err
	}

	if !e {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// Check if a given directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if !os.IsNotExist(err) {
		return false, err
	}

	return false, nil
}
