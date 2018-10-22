# gameserver

[![Build Status](https://travis-ci.org/the4thamigo-uk/gameserver.svg?branch=master)](https://travis-ci.org/the4thamigo-uk/gameserver?branch=master)
[![Coverage Status](https://coveralls.io/repos/the4thamigo-uk/gameserver/badge.svg?branch=master&service=github)](https://coveralls.io/github/the4thamigo-uk/gameserver?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/the4thamigo-uk/gameserver)](https://goreportcard.com/report/github.com/the4thamigo-uk/gameserver)


## Description

Gameserver is a simple implementation of a REST-based API that allows a user to play
simple games.

The game server currently only supports the game [Hangman](https://en.wikipedia.org/wiki/Hangman_(game)).
The game server currently uses an in-memory datastore to store game data.

The gameserver uses are a number of sub-packages the godocs can be found at
[/pkg/domain](https://godoc.org/github.com/the4thamigo-uk/gameserver/pkg/domain),
[/pkg/store](https://godoc.org/github.com/the4thamigo-uk/gameserver/pkg/store) and
[/pkg/server](https://godoc.org/github.com/the4thamigo-uk/gameserver/pkg/server)

## Getting Started

To build run

    go build ./cmd/...

Then to run an example server using english words run

    ./cmd/server/server --config ./config/english_config.yaml

From there you access the api from the root :

    curl -s http://127.0.0.1:8080

    {
      "_links": {
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "self": {
          "title": "Game server",
          "href": "/",
          "method": "GET"
        }
      }
    }

The hangman main index is given by :

    curl -s http://127.0.0.1:8080/hangman
    {
      "games": [],
      "_links": {
        "hangman:create": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        },
        "hangman:join": {
          "title": "Join hangman game",
          "href": "/hangman/{id}",
          "method": "GET"
        },
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "self": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        }
      }
    }

Create a new game with :

    curl -s -X POST http://127.0.0.1:8080/hangman/create
    {
      "game": {
        "id": {
          "id": "bf6qi6et5bvre3dv5t7g",
          "version": 1
        },
        "current": "      ",
        "turns": 6,
        "state": "play"
      },
      "_links": {
        "hangman:create": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        },
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "hangman:play:letter": {
          "title": "Guess a letter",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/1/letter/{letter}",
          "method": "PATCH"
        },
        "hangman:play:word": {
          "title": "Guess the word",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/1/word/{word}",
          "method": "PATCH"
        },
        "self": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        }
      }
    }

Guess a letter with :

    curl -s -X PATCH http://127.0.0.1:8080/hangman/bf6qi6et5bvre3dv5t7g/1/letter/a | jq
    {
      "game": {
        "id": {
          "id": "bf6qi6et5bvre3dv5t7g",
          "version": 2
        },
        "current": "  A   ",
        "turns": 6,
        "state": "play",
        "success": true
      },
      "_links": {
        "hangman:create": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        },
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "hangman:play:letter": {
          "title": "Guess a letter",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/2/letter/{letter}",
          "method": "PATCH"
        },
        "hangman:play:word": {
          "title": "Guess the word",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/2/word/{word}",
          "method": "PATCH"
        },
        "self": {
          "title": "Guess a letter",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/2/letter/{letter}",
          "method": "PATCH"
        }
      }
    }

Guess the word with :

    curl -s -X PATCH http://127.0.0.1:8080/hangman/bf6qi6et5bvre3dv5t7g/2/word/orange | jq
    {
      "game": {
        "id": {
          "id": "bf6qi6et5bvre3dv5t7g",
          "version": 3
        },
        "current": "ORANGE",
        "word": "ORANGE",
        "turns": 6,
        "state": "win",
        "success": true
      },
      "_links": {
        "hangman:create": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        },
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "self": {
          "title": "Guess the word",
          "href": "/hangman/bf6qi6et5bvre3dv5t7g/4/word/{word}",
          "method": "PATCH"
        }
      }
    }

List the games of hangman with :

    curl -s http://127.0.0.1:8080/hangman | jq
    {
      "games": [
        {
          "id": {
            "id": "bf6qi6et5bvre3dv5t7g",
            "version": 3
          },
          "current": "ORANGE",
          "word": "ORANGE",
          "turns": 6,
          "state": "win"
        }
      ],
      "_links": {
        "hangman:create": {
          "title": "Create hangman game",
          "href": "/hangman/create",
          "method": "POST"
        },
        "hangman:join": {
          "title": "Join hangman game",
          "href": "/hangman/{id}",
          "method": "GET"
        },
        "hangman:list": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        },
        "self": {
          "title": "Hangman game list",
          "href": "/hangman",
          "method": "GET"
        }
      }
    }

Alternatively use the test command-line client tool :

  ./cmd/client/client
