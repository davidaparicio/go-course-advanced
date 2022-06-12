#!/usr/bin/env bash

echo "Running some load (4 concurrent queries)"
curl http://localhost:8080 & \
echo "-------------------" & \
curl http://localhost:8080 & \
echo "-------------------" & \
curl -v http://localhost:8080 & \
echo "-------------------" & \
curl -v http://localhost:8080 & \
echo "-------------------" & \
sleep 10
