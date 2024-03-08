#!/bin/sh

set -e

INPUT_FILE=$1

ls
pwd

# build source code
CGO_ENABLED=0 GOOS=linux go build -o /app/circular-dependency-detector

/app/circular-dependency-detector $INPUT_FILE
# python3 ./visualizer/visualize.py
