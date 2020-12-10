#!/bin/sh

# Go to the source directory.
cd /app/src

# Start the headless Delve server.
# To find out more about delve flags take a look at: https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_debug.md
dlv debug --output=/app/mapi.debug --headless --listen=:6868 --log --api-version=2
