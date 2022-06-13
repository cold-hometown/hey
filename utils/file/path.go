package file

import (
	"errors"
	"os"
)

func GetHomeDirectoryPath() (string, error) {
	home := os.Getenv("HOME")

	if home != "" {
		return home, nil
	}

	return "", errors.New("Cannot find home directory. Let's check your environment.")

}
