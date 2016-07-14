Releasing
=========

 1. Update the `versionCode` and `versionName` in `app/build.gradle` to the next release version.
 2. Update `version` in `cli/main.go` to the next release version.
 3. Update the `CHANGELOG.md` for the impending release.
 3. `git commit -am "Prepare for release X.Y.Z."` (where X.Y.Z is the new version).
 4. `git tag -a X.Y.Z -m "Version X.Y.Z"` (where X.Y.Z is the new version).
 5. `git push && git push --tags`.
 6. Build the app `cd app && make build`.
 7. Build the cli `cd cli && make build`.
 8. Upload the binaries `app/build/outputs/apk/app-debug.apk`, `cli/bin/cli_linux_amd64`, `run.sh`, `start_emulator.sh` and `wait_for_emulator.sh`.
