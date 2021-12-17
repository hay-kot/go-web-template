package config

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v2"
	"github.com/ardanlabs/conf/v2/yaml"
	"io/ioutil"

	"os"
)

type Config struct {
	Web      WebConfig  `yaml:"web"`
	Database Database   `yaml:"database"`
	Log      LoggerConf `yaml:"logger"`
}

type WebConfig struct {
	Port string `yaml:"port" conf:"default:3000"`
	Host string `yaml:"host" conf:"default:127.0.0.1"`
}

// NewConfig parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func NewConfig(file string) (*Config, error) {
	var cfg Config

	const prefix = "API"

	var help string
	var err error
	if file != "" {
		yamlData, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		help, err = conf.Parse(prefix, &cfg, yaml.WithData(yamlData))
	} else {
		help, err = conf.Parse(prefix, &cfg)
	}

	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			// TODO: Evaluate if this should exit the program or return an error.
			os.Exit(0)
		}
		return &cfg, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
