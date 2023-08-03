package main

import (
	log "github.com/sirupsen/logrus"
)

// print hello world in a main function
func main() {
	log.Infof("Hello %s", "World")
}
