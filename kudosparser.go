package main

import (
	"regexp"
)

func parseNames(text string) []string {
	namePattern := regexp.MustCompile(`@[A-Za-z0-9]+`)
	return namePattern.FindAllString(text, -1)
}
