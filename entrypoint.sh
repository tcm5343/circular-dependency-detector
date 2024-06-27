#!/bin/sh

set -e

/app/circular-dependency-detector --input-file="$1" --fail-on-cycle="$2"
# python3 ./visualizer/visualize.py
