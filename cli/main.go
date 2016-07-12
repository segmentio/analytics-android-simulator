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
	usage   = `Analytics Android CLI

Usage:
  analytics track <event> [--properties=<properties>] [--context=<context>] [--writeKey=<writeKey>] [--userId=<userId>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics screen <name> [--properties=<properties>] [--context=<context>] [--writeKey=<writeKey>] [--userId=<userId>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics page <name> [--properties=<properties>] [--context=<context>] [--writeKey=<writeKey>] [--userId=<userId>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics identify [--traits=<traits>] [--context=<context>] [--writeKey=<writeKey>] [--userId=<userId>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics group --groupId=<groupId> [--traits=<traits>] [--properties=<properties>] [--context=<context>] [--writeKey=<writeKey>] [--userId=<userId>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics alias --userId=<userId> --previousId=<previousId> [--traits=<traits>] [--properties=<properties>] [--context=<context>] [--writeKey=<writeKey>] [--anonymousId=<anonymousId>] [--integrations=<integrations>] [--timestamp=<timestamp>]
  analytics -h | --help
  analytics --version

Options:
  -h --help     Show this screen.
  --version     Show version.`
)

func main() {
	log.SetHandler(cli.New(os.Stdout))
	log.Infof("arguments: %q", os.Args)

	arguments, err := docopt.Parse(usage, nil, true, version, false, false)
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

	log.Fatal("unknown command")
}

func track(arguments map[string]interface{}) {
	event := arguments["<event>"].(string)
	properties := getOptionalMap(arguments, "--properties")

	log.WithFields(log.Fields{
		"event":      event,
		"properties": properties,
	}).Info("simulating track call")

	var extras []string
	extras = append(extras, "-e", "type", "track")
	extras = append(extras, "-e", "event", event)
	for k, v := range properties {
		extras = append(extras, "-e", "properties_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(extras, arguments)
}

func screen(arguments map[string]interface{}) {
	name := arguments["<name>"].(string)
	properties := getOptionalMap(arguments, "--properties")

	log.WithFields(log.Fields{
		"name":       name,
		"properties": properties,
	}).Info("simulating screen call")

	var extras []string
	extras = append(extras, "-e", "type", "screen")
	extras = append(extras, "-e", "name", name)
	for k, v := range properties {
		extras = append(extras, "-e", "properties_"+k, fmt.Sprintf("%v", v))
	}
	runActivity(extras, arguments)
}

func getOptionalString(m map[string]interface{}, k string) string {
	v := m[k]
	if v == nil {
		return ""
	}
	return v.(string)
}

func getOptionalMap(m map[string]interface{}, k string) map[string]interface{} {
	rawMap := getOptionalString(m, k)
	if rawMap != "" {
		var parsedMap map[string]interface{}
		err := json.Unmarshal([]byte(rawMap), &parsedMap)
		if err != nil {
			log.WithError(err).WithField("key", k).Fatal("error parsing map")
		}
		return parsedMap
	}
	return make(map[string]interface{}, 0)
}

func runActivity(adbExtras []string, inputArgs map[string]interface{}) {
	adbArgs := []string{"shell", "am", "start", "-n", "com.segment.analytics.android.cli/com.segment.analytics.android.cli.MainActivity"}

	go tailLogcat()

	writeKey := getOptionalString(inputArgs, "--writeKey")
	if writeKey == "" {
		writeKey = os.Getenv("SEGMENT_WRITE_KEY")
		if writeKey == "" {
			log.Fatal("either $SEGMENT_WRITE_KEY or --writeKey must be provided")
		}
	}
	adbExtras = append(adbExtras, "-e", "writeKey", writeKey)

	log.Infof("running %v", append(adbArgs, adbExtras...))
	adb := exec.Command("adb", append(adbArgs, adbExtras...)...)
	out, err := adb.CombinedOutput()
	if err != nil {
		log.WithError(err).WithField("output", string(out)).Fatal("error running command")
	}
	log.Info(string(out))

	log.Info("waiting for command to run")
	time.Sleep(10 * time.Second)
	log.Info("waited for command to run")
}

func tailLogcat() {
	log.Info("tailing logcat")
	adb := exec.Command("adb", "logcat")
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
