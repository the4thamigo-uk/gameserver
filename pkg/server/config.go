package server

import (
	"github.com/miracl/conflate"
	"github.com/the4thamigo-uk/gameserver/pkg/domain"
)

// Config defines the main settings of the game server
type Config struct {
	Address string
	Hangman HangmanConfig
	// ... other games here ...
}

// HangmanConfig defines the main settings for the Hangman game
type HangmanConfig struct {
	Words []string
	Turns int
	dict  *domain.Dictionary
}

// LoadConfig loads a config from a local file or url
func LoadConfig(paths ...string) (*Config, error) {
	c, err := conflate.FromFiles(paths...)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = c.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	err = cfg.Hangman.init()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *HangmanConfig) init() error {
	cfg.dict = domain.NewDictionary(cfg.Words)
	return nil
}
