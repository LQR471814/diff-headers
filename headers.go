package main

import (
	"log"
	"strings"
)

type HeaderPair struct {
	Key   string
	Value string
}

type Headers struct {
	Pairs []HeaderPair
}

func (h *Headers) Get(key string) (string, bool) {
	for _, p := range h.Pairs {
		if p.Key == key {
			return p.Value, true
		}
	}
	return "", false
}

func ParseHeaders(raw string) Headers {
	pairs := []HeaderPair{}

	lines := strings.Split(raw, "\n")
	for _, l := range lines {
		if !strings.Contains(l, ":") {
			log.Println("Skipping line...", l)
			continue
		}
		split := strings.SplitN(l, ":", 2)
		pairs = append(pairs, HeaderPair{
			Key:   strings.ToLower(strings.Trim(split[0], " ")),
			Value: strings.ToLower(strings.Trim(split[1], " ")),
		})
	}

	return Headers{pairs}
}
