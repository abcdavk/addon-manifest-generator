#!/bin/bash
if [ "$1" = "" ]; then
  echo 'Usage generate-manifest --[option]'
  echo ' ' 
  echo '  OS                option'
  echo '  Windows amd64     --windows -w'
  echo '  Linux             --linux -l'
  echo ' '
fi
if [ "$1" = "--windows" -o "$1" = "-w" ]; then
  echo Building executable for windows amd64
  GOOS=windows GOARCH=amd64 go build -o build/generate-manifest.exe
fi
if [ "$1" = "--linux" -o "$1" = "-l" ]; then
  echo Building executable for linux
  go build -o build/generate-manifest
fi