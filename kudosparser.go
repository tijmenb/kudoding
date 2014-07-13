package main

import (
	"log"
	"regexp"
	"strconv"
)

func parseNames(text string) []string {
	namePattern := regexp.MustCompile(`@[A-Za-z0-9]+`)
	return namePattern.FindAllString(text, -1)
}

func parseAmount(text string) int {
	amountPattern := regexp.MustCompile(`[+-]\d+`)
	str := amountPattern.FindString(text)
	if str == "" {
		return 0
	}
	amount, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(amount)
}
