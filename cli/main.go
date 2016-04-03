package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/tj/docopt"
)

const (
	version = "0.1.0"
	usage   = `analytics-android-simulator.

Usage:
	sim track <event> [--properties=<p>]
	sim screen [--category=<c>] [--name=<n>] [--properties=<p>]
	sim identify [--userId=<id>] [--traits=<traits>]
	sim alias <userId>
	sim group <groupId> [--traits=<traits>]
	sim flush
	sim reset

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
	if arguments["screen"].(bool) {
		screen(arguments)
		return
	}
	if arguments["identify"].(bool) {
		identify(arguments)
		return
	}
	if arguments["alias"].(bool) {
		alias(arguments)
		return
	}
	if arguments["group"].(bool) {
		group(arguments)
		return
	}
	if arguments["flush"].(bool) {
		flush()
		return
	}
	if arguments["reset"].(bool) {
		reset()
		return
	}

	log.Fatal("unknown command")
}

func track(arguments map[string]interface{}) {
	event := arguments["<event>"].(string)
	var properties map[string]interface{}
	err := json.Unmarshal([]byte(arguments["--properties"].(string)), &properties)
	if err != nil {
		log.WithError(err).Fatal("invalid json")
	}

	log.WithFields(log.Fields{
		"event":      event,
		"properties": properties,
	}).Info("simulating track call")

	var args []string
	args = append(args, "-e", "type", "track")
	args = append(args, "-e", "event", event)
	for k, v := range properties {
		args = append(args, "-e", "properties_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(args)
}

func screen(arguments map[string]interface{}) {
	name := arguments["--name"].(string)
	category := arguments["--category"].(string)
	var properties map[string]interface{}
	err := json.Unmarshal([]byte(arguments["--properties"].(string)), &properties)
	if err != nil {
		log.WithError(err).Fatal("invalid json")
	}

	log.WithFields(log.Fields{
		"name":       name,
		"category":   category,
		"properties": properties,
	}).Info("simulating screen call")

	var args []string
	args = append(args, "-e", "type", "screen")
	args = append(args, "-e", "name", name)
	args = append(args, "-e", "category", category)
	for k, v := range properties {
		args = append(args, "-e", "properties_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(args)
}

func identify(arguments map[string]interface{}) {
	userID := arguments["--userId"].(string)
	var traits map[string]interface{}
	err := json.Unmarshal([]byte(arguments["--traits"].(string)), &traits)
	if err != nil {
		log.WithError(err).Fatal("invalid json")
	}

	log.WithFields(log.Fields{
		"userId": userID,
		"traits": traits,
	}).Info("simulating screen call")

	var args []string
	args = append(args, "-e", "type", "identify")
	args = append(args, "-e", "userId", userID)
	for k, v := range traits {
		args = append(args, "-e", "traits_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(args)
}

func alias(arguments map[string]interface{}) {
	userID := arguments["<userId>"].(string)

	log.WithFields(log.Fields{
		"userId": userID,
	}).Info("simulating alias call")

	var args []string
	args = append(args, "-e", "type", "alias")
	args = append(args, "-e", "userId", userID)
	runActivity(args)
}

func group(arguments map[string]interface{}) {
	userID := arguments["<userId>"].(string)
	var traits map[string]interface{}
	err := json.Unmarshal([]byte(arguments["--traits"].(string)), &traits)
	if err != nil {
		log.WithError(err).Fatal("invalid json")
	}

	log.WithFields(log.Fields{
		"userId": userID,
		"traits": traits,
	}).Info("simulating screen call")

	var args []string
	args = append(args, "-e", "type", "group")
	args = append(args, "-e", "userId", userID)
	for k, v := range traits {
		args = append(args, "-e", "traits_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(args)
}

func flush() {
	log.Info("simulating flush call")
	args := append([]string{}, "-e", "type", "flush")
	runActivity(args)
}

func reset() {
	log.Info("simulating reset call")
	args := append([]string{}, "-e", "type", "reset")
	runActivity(args)
}

func runActivity(cmdArgs []string) {
	args := []string{"shell", "am", "start", "-n", "com.segment.analyticsandroidsimulator/com.segment.analyticsandroidsimulator.MainActivity"}

	clearLogcat()
	go tailLogcat()

	adb := exec.Command("adb", append(args, cmdArgs...)...)
	out, err := adb.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(out)).Fatal("error running command")
	}
	log.Info(string(out))

	time.Sleep(2 * time.Second)
}

func clearLogcat() {
	// todo: fix, doesn't seem to be working.
	adb := exec.Command("adb", "logcat", "-c")
	out, err := adb.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(out)).Fatal("error clearing logcat")
	}
}

func tailLogcat() {
	adb := exec.Command("adb", "logcat", "-s", "Analytics")

	out, err := adb.StdoutPipe()
	if err != nil {
		log.WithError(err).Fatal("error reading adb logcat")
	}

	go func() {
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			log.Info(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}()

	err = adb.Run()
	if err != nil {
		log.WithError(err).Fatal("error running adb logcat")
	}
}
