#!/bin/bash
cd $(dirname $0)
cd ../

trap 'echo signal received!; kill $(jobs -p); wait;' SIGINT SIGTERM SIGUSR2 SIGQUIT

./cmd/main &

wait
