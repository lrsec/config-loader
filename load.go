// Package config
// Read configs from a json file

package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

const (
	RUN_MODE_ENV = "RUN_MODE"
)

var (
	configPath string
	runmode    string
)

func init() {
	var err error

	hasRunmode := false
	if runmode, hasRunmode = os.LookupEnv(RUN_MODE_ENV); !hasRunmode {
		panic("Can not find environment BILL_RUN_MODE")
	}

	workPath, err := os.Getwd()
	if err != nil {
		log.Error("Can not find workpath")
		panic(err)
	}

	configPath = filepath.Join(workPath, "conf", runmode)
	log.Info("Load config files under " + configPath)
}

func getRelativePath(fileName string) string {
	return filepath.Join(configPath, fileName)
}

// decode config from file. config must be a pointer
func LoadToml(config interface{}, confFile string) error {
	relativePath := getRelativePath(confFile)

	var err error

	_, err = toml.DecodeFile(relativePath, config)
	if err != nil {
		log.Infof("decode config file in relative path: %s error. try absolute path. error: %v", err)

		_, err = toml.DecodeFile(confFile, config)
	}

	return err
}
