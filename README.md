# analytics-android-simulator

App that lets you simulate sending arbitrary Segment events.

# Building

1. `cd app && make build`.

2. `cd cli && make build`.

3. `docker build -t analytics-android-simulator .`

# Structure

Provides 2 components:

1. App: The actual Android app that runs on a device.

2. CLI: CLI app to trigger events.

The CLI parses the input, converts them to intent extras and launches the sample app with the intent.
The sample app reads the extras and makes the analytics call.
