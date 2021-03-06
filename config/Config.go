package config

import (
	"encoding/json"
	"github.com/scinna/CLIent/utils"
	"io/ioutil"
	"os"
)

var CopyToClipboard = false

type Config struct {
	ServerURL string
	Username  string
	Token     string

	DefaultTitle       StringCommand
	DefaultDescription StringCommand
	DefaultVisibility  utils.Visibility
	DefaultCollection  string
}

func (cfg Config) IsConfigValid() bool {
	return len(cfg.ServerURL) > 0 && len(cfg.Username) > 0 && len(cfg.Token) > 0
}

func (cfg Config) AreDefaultValid() bool {
	return len(cfg.DefaultTitle) > 0 && len(cfg.DefaultVisibility.ToString()) > 0// && len(cfg.DefaultCollection) > 0
}

// WriteConfiguration writes the config to it's default path
func WriteConfiguration(cfg *Config) error {
	content, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	// 0600 = Only the owner can read it
	err = ioutil.WriteFile(getConfigFile(), content, 0600)
	if err != nil {
		return err
	}

	return nil
}

// ReadConfiguration reads the configuration from it's default path. The boolean is false when the configuration was not found
func ReadConfiguration() (*Config, error) {
	cfgFile := getConfigFile()

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		err = createDefaultConfiguration()
		if err != nil {
			return &Config{}, err
		}
	}

	file, err := os.Open(cfgFile)
	if err != nil {
		panic("Can't open the file!")
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Can't read the file (Does it have correct permissions?)")
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)

	if err != nil {
		panic("The configuration file has some errors in it. Please remove it and start again.")
	}

	return &cfg, err
}
