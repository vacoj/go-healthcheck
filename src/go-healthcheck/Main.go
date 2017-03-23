package main

import (
	"fmt"
	"time"
)

func showVersion() {
	name := "LB-Toggle " + VERSION
	fmt.Println(name)
}

func main() {
	SETTINGS.parseSettingsFile()

	// Monitor application for health status
	STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	if SETTINGS.State.StartupState {
		// give the targets a bit to catch up
		time.Sleep(time.Duration(len(STATUS.Targets)) * time.Second)
		if STATUS.isOk() {
			STATUS.State.OK = true
		}
	}

	WG.Wait()
}
