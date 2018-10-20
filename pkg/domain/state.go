package domain

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type State int

const (
	Play State = iota + 1
	Win
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

func (s State) String() string {
	return stateToString[s]
}

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

func (s *State) MarshalJSON() ([]byte, error) {
	val, ok := stateToString[*s]
	if !ok {
		return nil, errors.New("Invalid state value")
	}
	return json.Marshal(val)
}
