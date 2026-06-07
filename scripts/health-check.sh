#!/bin/bash

curl -f "$1/health"

if [ $? -eq 0 ]; then
  echo "Health check passed"
else
  echo "Health check failed"
  exit 1
fi
