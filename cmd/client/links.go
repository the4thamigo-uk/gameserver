package main

type link struct {
	Href   string
	Method string
	Title  string
}

type links map[string]*link
