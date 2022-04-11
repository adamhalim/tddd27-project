package chunk

import (
	"os"
)

func fileExists(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return nil
	}
	// We don't care if it's a directory or file,
	// return error anyways
	return err
}
