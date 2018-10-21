package main

import (
	"github.com/the4thamigo-uk/gameserver/pkg/domain"
)

type errResponse struct {
	Message string
	Code    int
}

type response struct {
	Error *errResponse
	Game  *domain.HangmanResult
	Games []*domain.HangmanResult
	Links links `json:"_links"`
}
