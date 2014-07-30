package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func parseInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parseSlackTimestamp(timestamp string) time.Time {
	parts := strings.Split(timestamp, ".")
	sec := parseInt(parts[0])
	nsec := parseInt(parts[1])
	return time.Unix(sec, nsec)
}
