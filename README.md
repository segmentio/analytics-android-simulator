# analytics-android-simulator

App that lets you simulate sending arbitrary Segment events.

Provides 2 components:
1. App: The actual Android app that runs on a device.
2. CLI: CLI app to trigger events.

The app and cli communicate via web sockets. The cli accepts json as input, and
sends it to the app which parses it and calls the analytics SDK.
