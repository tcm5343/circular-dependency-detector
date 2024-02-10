#!/bin/sh

set -e

INPUT_FILE=$1

/app/circular-dependency-detector $INPUT_FILE
# python3 ./visualizer/visualize.py
