package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Screen   Screen   `yaml:"screen"`
	Ticktick Ticktick `yaml:"ticktick" mapstructure:"ticktick"`
	Todoist  Todoist  `yaml:"todoist" mapstructure:"todoist"`
}

type Screen struct {
	Output string `yaml:"output"`
}

type Ticktick struct {
	ApiUrl       string `yaml:"api_url" mapstructure:"api_url"`
	AccessToken  string `yaml:"access_token" mapstructure:"access_token"`
	FocusProject string `yaml:"focus_project" mapstructure:"focus_project"`
}

type Todoist struct {
	ApiUrl      string   `yaml:"api_url" mapstructure:"api_url"`
	AccessToken string   `yaml:"access_token" mapstructure:"access_token"`
	Project     string   `yaml:"project"`
	Labels      []string `yaml:"labels"`
}

func LoadConfig(files ...string) (*Config, error) {
	if files != nil && len(files) > 0 {
		for _, file := range files {
			if _, err := os.Stat(file); err == nil {
				viper.SetConfigFile(file)
				err := viper.MergeInConfig()
				if err != nil { // Handle errors reading the config file
					panic(fmt.Errorf("Fatal error config file: %w \n", err))
				}
			} else {
				fmt.Printf("Could not load config file: %s", file)
			}
		}
	}

	viper.BindEnv(
		"SCREEN.OUTPUT",
		"TICKTICK.USERNAME",
		"TICKTICK.PASSWORD",
		"TICKTICK.ACCESS_TOKEN",
		"TICKTICK.FOCUS_PROJECT",
		"TODOIST.ACCESS_TOKEN",
		"TODOIST.PROJECT",
		"TODOIST.LABELS",
	)
	viper.AutomaticEnv()

	if err := viper.MergeInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	cfg := Config{}
	err := viper.Unmarshal(&cfg)

	return &cfg, err
}

// Save config to file
func (c *Config) Save(file string) error {
	yfile, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, yfile, 0644)
	if err != nil {
		return err
	}
	return nil
}
