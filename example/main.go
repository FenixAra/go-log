package main

import (
	"errors"

	"github.com/FenixAra/go-log/log"
)

func main() {
	config := log.NewConfig("Test_App")
	l := log.New(config)

	l.Fatalf("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("No rows"))
}
