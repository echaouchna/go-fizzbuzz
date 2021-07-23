package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting fizzbuzz server...")
	statsChannel = make(chan FizzbuzzOptions)
	defer close(statsChannel)
	go collectStats()
	log.Infof("Listening on port 10000")
	handleRequests()
}
