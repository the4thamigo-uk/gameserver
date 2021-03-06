package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jtacoma/uritemplates"
	"github.com/the4thamigo-uk/gameserver/pkg/domain"
	"io/ioutil"
	"net/http"
	pkgurl "net/url"
	"strconv"
)

var relOrder = []string{"index",
	"hangman:list",
	"hangman:join",
	"hangman:create",
	"hangman:play:letter",
	"hangman:play:word",
}

func run(url string) error {
	rel := "index"
	rsp, err := request("GET", url, nil)
	if err != nil {
		return err
	}
	for {
		rel2, rsp2, err := processResponse(url, rel, rsp)
		if err != nil {
			return err
		}
		displayResult(rel2, rsp2)
		if rsp2.Error != nil {
			// response contained an error so retry processing of last message
			continue
		}
		rel, rsp = rel2, rsp2
	}
}

func processResponse(rootURL string, rel string, rsp *response) (string, *response, error) {
	fmt.Println(rsp.Links["self"].Title)

	opts := optionsFromLinks(rsp.Links)

	opt, err := enterOption(opts)
	if err != nil {
		return "", nil, err
	}

	path, err := enterValues(opt)
	if err != nil {
		return "", nil, err
	}

	url, err := joinURL(rootURL, path)
	if err != nil {
		return "", nil, err
	}

	rsp, err = request(opt.link.Method, url, nil)
	if err != nil {
		return "", nil, err
	}

	return opt.rel, rsp, nil
}

func enterOption(opts options) (*option, error) {
	for {
		displayOptions(opts)
		fmt.Println("Please choose an option :")
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(s)
		if err == nil && i >= 1 && i <= len(opts) {
			return opts[i-1], nil
		}
		fmt.Println("Option not valid. Please try again.")
	}
}

func displayOptions(opts options) {
	for i, opt := range opts {
		fmt.Printf("%v) %v\n", i+1, opt.link.Title)
	}
}

func enterValues(opt *option) (string, error) {
	tmpl, err := uritemplates.Parse(opt.link.Href)
	if err != nil {
		return "", err
	}
	vals := map[string]interface{}{}
	for _, n := range tmpl.Names() {
		for {
			fmt.Printf("Please enter a value for '%v':\n", n)
			var val string
			_, err := fmt.Scanln(&val)
			if err != nil {
				return "", err
			}
			if len(val) > 0 {
				vals[n] = val
				break
			}
			fmt.Println("Value cannot be blank. Please try again.")
		}
	}

	return tmpl.Expand(vals)
}

func displayResult(rel string, rsp *response) {
	if rsp.Error != nil {
		fmt.Printf("The server reported an error : '%v'\n", rsp.Error.Message)
	} else if rsp.Game != nil {
		if rsp.Game.Success != nil {
			if *rsp.Game.Success {
				fmt.Println("Your guess was CORRECT.")
			} else {
				fmt.Println("Your guess was INCORRECT.")
			}
		}
		switch rsp.Game.State {
		case domain.Win:
			fmt.Printf("You have WON with %v turn(s) remaining. The word was '%v'.\n", rsp.Game.Turns, *rsp.Game.Word)
		case domain.Lose:
			fmt.Printf("You have LOST. The word was '%v'.\n", *rsp.Game.Word)
		case domain.Play:
			fmt.Printf("The word is '%v' and you have %v turn(s) remaining.\n", rsp.Game.Current, rsp.Game.Turns)
		}
	} else if rsp.Games != nil {
		if len(rsp.Games) == 0 {
			fmt.Printf("There are no games on the server.\n")
		} else {
			fmt.Printf("Games on server :\n")
			for i, g := range rsp.Games {
				fmt.Printf("%v). Game '%v' is at version '%v', with %v turn(s) remaining. Current word is '%v' and the game state is '%v'.\n", i+1, g.ID.ID, g.ID.Version, g.Turns, g.Current, g.State)
			}
		}
	}
}

func request(method string, url string, body []byte) (*response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return parseBody(rsp)
}

func parseBody(r *http.Response) (*response, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var rsp response
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

func joinURL(base string, path string) (string, error) {
	p, err := pkgurl.Parse(path)
	if err != nil {
		return "", err
	}
	b, err := pkgurl.Parse(base)
	if err != nil {
		return "", err
	}
	url := b.ResolveReference(p).String()
	return url, nil
}
