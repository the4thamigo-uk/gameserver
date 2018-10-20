package server

import (
	"github.com/miracl/conflate"
	"github.com/the4thamigo_uk/gameserver/pkg/domain"
)

type Config struct {
	Address string
	Hangman HangmanConfig
	// ... other games here ...
}

type HangmanConfig struct {
	Words []string
	Turns int
	dict  *domain.Dictionary
}

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
	err = cfg.Hangman.Init()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *HangmanConfig) Init() error {
	cfg.dict = domain.NewDictionary(cfg.Words)
	return nil
}
