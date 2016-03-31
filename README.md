# analytics-android-simulator

App that lets you simulate sending arbitrary Segment events.

Provides 2 components:
1. App: The actual Android app that runs on a device.
2. CLI: CLI app to trigger events.

The CLI parses the input, converts them to intent extras and launches the sample app with the intent.
The sample app reads the extras and makes the call.
