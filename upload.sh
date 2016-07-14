#!/bin/bash

github-release upload \
    --user segmentio \
    --repo analytics-android-simulator \
    --tag $1 \
    --name $2 \
    --label cli.sh \
    --file cli/bin/cli_linux_amd64

github-release upload \
    --user segmentio \
    --repo analytics-android-simulator \
    --tag $1 \
    --name $2 \
    --label app.apk \
    --file app/build/outputs/apk/app-debug.apk

github-release upload \
    --user segmentio \
    --repo analytics-android-simulator \
    --tag $1 \
    --name $2 \
    --label run.sh \
    --file run.sh

github-release upload \
    --user segmentio \
    --repo analytics-android-simulator \
    --tag $1 \
    --name $2 \
    --label start_emulator.sh \
    --file start_emulator.sh

github-release upload \
    --user segmentio \
    --repo analytics-android-simulator \
    --tag $1 \
    --name $2 \
    --label wait_for_emulator.sh \
    --file wait_for_emulator.sh
