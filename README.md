# analytics-android-simulator

App that lets you simulate sending arbitrary Segment events.

Provides 2 components:
1. App: The actual Android app that runs on a device.
2. CLI: CLI app to trigger events.

The CLI parses the input, converts them to intent extras and launches the sample app with the intent.
The sample app reads the extras and makes the call.

# Usage

Start an emulator and install the app (easiest via Android Studio).

`go run cli/main.go track foo --properties="{\"foo\": 23.1}"`

```
analytics-android-simulator.

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
  --version             Show version.
```
