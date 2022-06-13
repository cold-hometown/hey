package config

import (
	"github.com/cold-hometown/hey/utils/file"
)

type Meta struct {
	path string
	name string
}

const defaultFileName = "hey.yaml"

func (m Meta) SetPath(path string) {
	m.path = path
}

func (m Meta) GetPath() (string, error) {
	if m.path != "" {
		return m.path, nil
	}

	homedir, err := file.GetHomeDirectoryPath()

	if err != nil {
		return "", err
	}

	return homedir, nil
}

func (m Meta) SetName(name string) {
	m.name = name
}

func (m Meta) GetName() string {
	if m.name != "" {
		return m.name
	}

	return defaultFileName
}
