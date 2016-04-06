# analytics-android-simulator

App that lets you simulate sending arbitrary Segment events.

# One-Time Setup

Clone the repo.

One-Time: Run `docker build -t analytics-android-simulator .`

# Usage

Everything runs inside the docker container, so you need to run all the simulator commands inside the container too (for now).

`docker run -it analytics-android-simulator /bin/bash`

Once you're in, just use the simulator CLI!

# CLI Component

`simulator track foo --properties="{\"foo\": 23.1}"`

```
analytics-android-simulator.

Usage:
  simulator track <event> [--properties=<p>]
  simulator screen [--category=<c>] [--name=<n>] [--properties=<p>]
  simulator identify [--userId=<id>] [--traits=<traits>]
  simulator alias <userId>
  simulator group <groupId> [--traits=<traits>]
  simulator flush
  simulator reset

  simulator -h | --help
  simulator --version

Options:
  -h --help             Show this screen.
  --version             Show version.
```

# Building

1. `cd app && ./gradlew build`.

2. Copy `app/app/build/outputs/apk/app-debug.apk` to `bin`.

3. `go build -o bin/simulator ./...`.

4. `docker build -t analytics-android-simulator .`

# Structure

Provides 2 components:

1. App: The actual Android app that runs on a device.

2. CLI: CLI app to trigger events.

The CLI parses the input, converts them to intent extras and launches the sample app with the intent.
The sample app reads the extras and makes the call.
