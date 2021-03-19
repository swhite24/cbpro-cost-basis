package config

import (
	"errors"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	format = "2006-01-02"
)

type (
	// Config defines configuration for purchasing from coinbase pro
	Config struct {
		// Client configuration
		BaseURL    string `json:"base_url"`
		Key        string `json:"key" mapstructure:"key"`
		Passphrase string `json:"passphrase" mapstructure:"passphrase"`
		Secret     string `json:"secret" mapstructure:"secret"`

		// Purchase configuration
		Product      string `json:"product" mapstructure:"product"`
		StartDateStr string `json:"start" mapstructure:"start"`
		EndDateStr   string `json:"end" mapstructure:"end"`
		StartDate    time.Time
		EndDate      time.Time
	}
)

// InitializeConfig delivers the initialized config
func InitializeConfig(flags *pflag.FlagSet) (*Config, error) {
	var err error
	viper.BindPFlags(flags)

	viper.SetEnvPrefix("CBPRO_COST_BASIS")
	viper.AutomaticEnv()

	c := Config{BaseURL: "https://api.pro.coinbase.com"}
	if err = viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	if c.StartDateStr == "" {
		return nil, errors.New("start is required")
	}
	if err = c.parseDates(); err != nil {
		return nil, errors.New("Unable to parse time range")
	}

	return &c, nil
}

func (c *Config) parseDates() error {
	var err error
	if c.StartDate, err = time.Parse(format, c.StartDateStr); err != nil {
		return err
	}
	if c.EndDate, err = time.Parse(format, c.EndDateStr); err != nil {
		c.EndDate = time.Now()
	}
	return nil
}
