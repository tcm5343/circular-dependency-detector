#!/bin/sh

set -e

echo "Hello $1"

pwd

/app/circular-dependency-detector
# python3 ./visualizer/visualize.py
