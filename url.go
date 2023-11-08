package main

import (
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(?:(?:https?|ftp)://)?(?:www\.)?[a-zA-Z0-9-]+\.[a-zA-Z]+(?::\d+)?(?:/[^/\s$]*)*$`)

type Url struct {
	Url   string
	Error string
}

func (url *Url) Validate() bool {
	match := urlRegex.Match([]byte(url.Url))
	if match == false {
		url.Error = "Invalid URL"
	}

	return len(url.Error) == 0
}
