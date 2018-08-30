package logger

import (
	"os"
)

func createDir(path string) error {
	if _, errIsNotExist := os.Stat(path); os.IsNotExist(errIsNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
