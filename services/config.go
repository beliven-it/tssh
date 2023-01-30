package services

import (
	"crypto/md5"
	"os"
	"path/filepath"

	"tssh/defs"
	"tssh/templates"

	"github.com/spf13/viper"
)

type config struct {
}

type Config interface {
	IsInitialized() bool
	Load() error
	Init() error
}

func (s *config) IsInitialized() bool {
	h := md5.New()
	file, err := os.ReadFile(defs.ConfigFilePath)
	if err != nil {
		return false
	}

	return string(h.Sum(file)) != string(h.Sum([]byte(templates.Config)))
}

func (s *config) Load() error {
	viper.SetConfigFile(defs.ConfigFilePath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func (s *config) Init() error {
	if s.IsInitialized() {
		return nil
	}

	err := os.MkdirAll(filepath.Dir(defs.ConfigFilePath), 0755)
	if err != nil {
		return err
	}

	_, err = os.Stat(defs.ConfigFilePath)
	if err == nil {
		return nil
	} else if os.IsNotExist(err) {
		err := os.WriteFile(defs.ConfigFilePath, []byte(templates.Config), 0755)
		if err != nil {
			return err
		}

		s.Load()
		return nil
	} else {
		return err
	}
}

func NewConfigService() Config {
	return &config{}
}
