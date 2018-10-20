package domain

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// State is an enum representing the state of a game.
type State int

const (
	// Play indicates that gameplay is in progress,
	Play State = iota + 1
	// Win indicates that the game is complete and the client has won.
	Win
	// Lose indicates that the game is complete and the client has lost.
	Lose
)

var (
	stateToString = map[State]string{
		Play: "play",
		Win:  "win",
		Lose: "lose",
	}
	stringToState = map[string]State{}
)

func init() {
	for k, v := range stateToString {
		stringToState[v] = k
	}
}

// String returns the string representation of the enum (play,win.lose)
func (s State) String() string {
	return stateToString[s]
}

// UnmarshalJSON unmarshals JSON into an instance of the State enum.
func (s *State) UnmarshalJSON(b []byte) error {
	var data string
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	val, ok := stringToState[data]
	if !ok {
		return errors.New("Invalid state value")
	}
	*s = val
	return nil
}

// MarshalJSON marshals an instance of State enum into JSON.
func (s *State) MarshalJSON() ([]byte, error) {
	val, ok := stateToString[*s]
	if !ok {
		return nil, errors.New("Invalid state value")
	}
	return json.Marshal(val)
}
