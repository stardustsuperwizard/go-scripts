package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}


	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcomes!", name)
	return message, nil
}