#!/usr/bin/env bash

start_server() {
    echo "Start server ..."
    CGO_ENABLED=1
    go run -race main.go &
    server_pid=$!
    sleep 3
    my_pid=$(lsof -t -i tcp:8080)
    # Wait for the server to start (max 10 seconds)
    # for attempt in {1..10}; do
    #     echo "Attempt: ${attempt}"
    #     my_pid=$(lsof -t -i tcp:8080)

    #     if [[ -n $my_pid ]]; then
    #         # Make sure the running server is the one we just started.
    #         if [[ $my_pid -ne $server_pid ]]; then
    #             echo "ERROR: Multiple Servers running."
    #             echo "→ lsof -t -i tcp:8080 | xargs kill"
    #             exit 1
    #         fi
    #         break
    #     fi
    #     sleep 1
    # done

    if [[ -z $my_pid ]]; then
        echo "ERROR: Timeout while waiting for Server"
        exit 1
    fi
}

stop_server() {
    echo "Stop Server ..."
    kill $server_pid
    kill $my_pid
}

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 

start_server
curl http://localhost:8080
curl http://localhost:8080
hey -n 1000 http://localhost:8080
curl http://localhost:8080
stop_server
# https://stackoverflow.com/a/63982932