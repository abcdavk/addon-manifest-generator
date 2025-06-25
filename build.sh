#!/bin/bash
if [ "$1" = "" ]; then
  echo "Usage 'generate-manifest --[option]'"
  echo '  OS             option'
  echo '  All            --all -a'
  echo '  Windows amd64  --windows -w'
  echo '  Linux          --linux -l'
fi
if [ "$1" = "--all" -o "$1" = "-a" ]; then
  echo Building executable for windows amd64 and linux
  go build -C src/ -o ../build/generate-manifest-linux
  GOOS=windows GOARCH=amd64 go build -C src/ -o ../build/generate-manifest-windows.exe
  echo '  linux: build/generate-manifest'
  echo '  windows: build/generate-manifest.exe'
fi
if [ "$1" = "--windows" -o "$1" = "-w" ]; then
  echo Building executable for windows amd64
  GOOS=windows GOARCH=amd64 go build -C src/ -o ../build/generate-manifest-windows.exe
  echo '  windows: build/generate-manifest.exe'
fi
if [ "$1" = "--linux" -o "$1" = "-l" ]; then
  echo Building executable for linux
  go build -C src/ -o ../build/generate-manifest-linux
  echo '  linux: build/generate-manifest'
fi