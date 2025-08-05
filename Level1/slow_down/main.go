package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	splitMsg := strings.Split(msg, " ")
	for _, msg := range splitMsg {
		splitLetter := strings.Split(msg, "")
		var letters string
		for i, letter := range splitLetter {
			letters += strings.Repeat(letter, i+1)
		}
		print(letters)
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
