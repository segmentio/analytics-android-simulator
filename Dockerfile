FROM jacekmarchwicki/android

# Create and Launch Emulator
RUN emulator -avd test -no-audio -no-window -force-32bit &

# Copy and install App
COPY bin/app-debug.apk /app-debug.apk
RUN adb install app-debug.apk

# Copy CLI
ADD bin/simulator /usr/local/bin/
