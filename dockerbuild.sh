#!/bin/sh
set -e
cd /TalesFromTheRim/
./configure
cd /TalesFromTheRim/src
make
cd /TalesFromTheRim/
./autorun
