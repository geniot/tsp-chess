#!/bin/bash

pkill -f hwt

if [ "$#" -gt 0 ]; then
 /mnt/SDCARD/Apps/Chess/chess "$@"
else
  progdir=$(dirname "$0")
  cd $progdir
  ./chess
fi