package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	MySQL struct {
		DSN string `yaml:"dsn"`
	} `yaml:"mysql"`
	CampaignRefreshIntervalMinutes int `yaml:"campaign_refresh_interval_minutes"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	dec := yaml.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) CampaignRefreshInterval() time.Duration {
	if c.CampaignRefreshIntervalMinutes <= 0 {
		return 5 * time.Minute
	}
	return time.Duration(c.CampaignRefreshIntervalMinutes) * time.Minute
}
