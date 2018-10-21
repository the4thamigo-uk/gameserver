package main

import (
	"github.com/the4thamigo-uk/gameserver/pkg/domain"
)

type response struct {
	Game  *domain.HangmanResult
	Games []*domain.HangmanResult
	Links links `json:"_links"`
}
