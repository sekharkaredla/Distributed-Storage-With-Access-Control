package main

import (
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

func main() {
	// From now on the code should write print statements like this so that logs can easily studied
	log.Info.Println("for info, if you want to give some info")
	log.Trace.Println("for tracing purposes")
	log.Warning.Println("for warning purposes")
	log.Error.Println("for error purposes")
}
