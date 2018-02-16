package config

import (
	"strings"
	"testing"
)

func TestGetRelativePath(t *testing.T) {
	originPath := "test"
	path := getRelativePath(originPath)

	if !strings.HasSuffix(path, originPath) {
		t.Errorf("getRelativePath do not have correct result. origin: %v. result: %v", originPath, path)
	}
}

type ServerCfg struct {
	Url  string `toml:"url"`
	Port int32  `toml:"port"`
}

type TestCfg struct {
	Id     int32     `toml:"id"`
	Desc   string    `toml:"desc"`
	Server ServerCfg `toml:"server"`
}

func TestLoad(t *testing.T) {
	configFile := "testconfig.toml"
	cfg := &TestCfg{}

	if err := LoadToml(cfg, configFile); err != nil {
		t.Errorf("Load error: %v", err)
	}

	if cfg.Id != 123 || cfg.Desc != "test config" || cfg.Server.Url != "http://test.test" || cfg.Server.Port != 8080 {
		t.Errorf("content is not correct. cfg: %+v", cfg)
	}

}
