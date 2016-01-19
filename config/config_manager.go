// Package config
// Read configs from a json file

package config

import (
	"encoding/json"
	errors "github.com/lrsec/errors/wrapper"
	"os"
)

func LoadConfig(fileName string, config interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return errors.New(err)
	}

	return nil
}
