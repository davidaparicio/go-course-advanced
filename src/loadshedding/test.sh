#!/usr/bin/env bash

echo "Running some load (4 concurrent queries)"
curl http://localhost:8080 & \
curl http://localhost:8080 & \
curl -v http://localhost:8080 & \
curl -v http://localhost:8080 & \
sleep 10