package server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/the4thamigo_uk/gameserver/pkg/domain"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testConfig() *Config {
	return &Config{
		Address: ":8080",
		Hangman: HangmanConfig{
			Words: []string{"word"},
			Turns: 6,
			dict:  domain.NewDictionary([]string{"word"}),
		},
	}
}

func TestServer_Start(t *testing.T) {

	s := NewServer(testConfig())
	var err error
	go func() {
		err = s.ListenAndServe()
	}()
	assert.Nil(t, err)
	s.Shutdown()
}

func unmarshalReader(r io.Reader, obj interface{}) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, json.Unmarshal(b, obj)
}

func testServerRequest(s *Server, method string, url string, rsp interface{}) ([]byte, error) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	s.serveHTTP(w, r)
	b, err := unmarshalReader(w.Result().Body, &rsp)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type testHangmanResponse struct {
	State domain.HangmanResult
	Links map[string]Link `json:"_links"`
}

func TestServer_PlayLetterToWin(t *testing.T) {
	s := NewServer(testConfig())
	var r1 testHangmanResponse
	b1, err := testServerRequest(s, "POST", "/hangman/create", &r1)
	assert.Nil(t, err)
	t.Log(string(b1))
	assert.Equal(t, 1, r1.State.ID.Version)
	assert.Equal(t, "    ", r1.State.Current)
	assert.Equal(t, hangmanDefaultTurns, r1.State.Turns)
	assert.Equal(t, domain.Play, r1.State.State)
	//assert.Equal(t, "", r1.State.Used)

	// play valid letter
	link2 := r1.Links[relHangmanPlayLetter]
	var r2 testHangmanResponse
	url2 := strings.Replace(link2.Href, "{letter}", "w", -1)
	b2, err := testServerRequest(s, link2.Method, url2, &r2)
	assert.Nil(t, err)
	t.Log(string(b2))

	assert.Equal(t, r1.State.ID.ID, r2.State.ID.ID)
	assert.Equal(t, 2, r2.State.ID.Version)
	assert.Equal(t, "W   ", r2.State.Current)
	assert.Equal(t, hangmanDefaultTurns, r2.State.Turns)
	assert.Equal(t, domain.Play, r2.State.State)
	assert.True(t, r2.State.Success)
	//assert.Equal(t, "W", r2.State.Used)

	link3 := r2.Links[relHangmanPlayLetter]
	var r3 testHangmanResponse
	url3 := strings.Replace(link3.Href, "{letter}", "x", -1)
	b3, err := testServerRequest(s, link3.Method, url3, &r3)
	assert.Nil(t, err)
	t.Log(string(b3))

	assert.Equal(t, r1.State.ID.ID, r3.State.ID.ID)
	assert.Equal(t, 3, r3.State.ID.Version)
	assert.Equal(t, "W   ", r3.State.Current)
	assert.Equal(t, hangmanDefaultTurns-1, r3.State.Turns)
	assert.Equal(t, domain.Play, r3.State.State)
	assert.False(t, r3.State.Success)
	//assert.Equal(t, "W", r3.State.Used)
}
