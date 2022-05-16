package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
		return "", errors.New("empty name")
	}
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}