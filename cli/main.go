package main

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/tj/docopt"
)

const (
	version = "0.1.0"
	usage   = `analytics-android-simulator.

Usage:
  sim track <event> [--properties=<props>]
  sim -h | --help
  sim --version

Options:
  -h --help             Show this screen.
  --version             Show version.`
)

func main() {
	log.SetHandler(cli.New(os.Stdout))

	arguments, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.WithError(err).Fatal("invalid arguments")
	}

	if arguments["track"].(bool) {
		track(arguments)
		return
	}
}

func track(arguments map[string]interface{}) {
	event := arguments["<event>"].(string)
	props := arguments["--properties"].(string)
	var properties map[string]interface{}
	err := json.Unmarshal([]byte(props), &properties)
	if err != nil {
		log.WithError(err).Fatal("invalid json")
	}

	log.WithFields(log.Fields{
		"event":      event,
		"properties": properties,
	}).Info("simulating track call")

	cmd := exec.Command("adb", "shell", "am", "start", "-n", "com.segment.analyticsandroidsimulator/com.segment.analyticsandroidsimulator.MainActivity")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(out)).Fatal("error running command")
	}
}
