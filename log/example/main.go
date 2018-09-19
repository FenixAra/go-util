package main

import (
	"errors"

	"github.com/FenixAra/go-util/log"
)

func main() {
	config := log.NewConfig("", "Debug", "Short", "Test_App", "", "", "")
	l := log.New(config)

	l.Fatalf("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("No rows"))
}
