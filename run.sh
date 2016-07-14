#!/bin/bash

set -ex

./start_emulator.sh&

sleep 5

./wait_for_emulator.sh

adb install app.apk

./cli.sh "$@"
