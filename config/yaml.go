package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func InitConfig(filename string, c interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	return yaml.NewDecoder(file).Decode(c)
}
