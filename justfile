#!/usr/bin/env just --justfile

hello:
  echo "hello world"

run-go-by-example:
    echo "Arrays\n"
    cd src/github.com/pawelhenek/go-by-example && go run ./arrays.go

    echo "Atomitc counters"
    cd src/github.com/pawelhenek/go-by-example && go run ./atomic_counters.go
    cd src/github.com/pawelhenek/go-by-example && go run ./base64-encoding.go
    cd src/github.com/pawelhenek/go-by-example && go run ./channel_buffering.go
    cd src/github.com/pawelhenek/go-by-example && go run ./channel_directions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./channel_synchronization.go
    cd src/github.com/pawelhenek/go-by-example && go run ./channels.go
    cd src/github.com/pawelhenek/go-by-example && go run ./closing_channels.go
    cd src/github.com/pawelhenek/go-by-example && go run ./closures.go
    cd src/github.com/pawelhenek/go-by-example && go run ./collection-functions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./command-line-arguments.go pierwszy drugi trzeci
    cd src/github.com/pawelhenek/go-by-example && go run ./command-line-flags.go
    cd src/github.com/pawelhenek/go-by-example && go run ./constants.go
    cd src/github.com/pawelhenek/go-by-example && go run ./defer.go
    cd src/github.com/pawelhenek/go-by-example && go run ./environment-variables.go
    cd src/github.com/pawelhenek/go-by-example && go run ./epoch.go
    cd src/github.com/pawelhenek/go-by-example && go run ./errors.go
    cd src/github.com/pawelhenek/go-by-example && go run ./execing-processes.go
    cd src/github.com/pawelhenek/go-by-example && go run ./for.go
    cd src/github.com/pawelhenek/go-by-example && go run ./functions.go
    #cd src/github.com/pawelhenek/go-by-example && go run ./goroutines.go
    cd src/github.com/pawelhenek/go-by-example && go run ./hello.go
    cd src/github.com/pawelhenek/go-by-example && go run ./if-else.go
    cd src/github.com/pawelhenek/go-by-example && go run ./interafces.go
    cd src/github.com/pawelhenek/go-by-example && go run ./json.go
    #cd src/github.com/pawelhenek/go-by-example && go run ./line-filters.go
    cd src/github.com/pawelhenek/go-by-example && go run ./maps.go
    cd src/github.com/pawelhenek/go-by-example && go run ./methods.go
    cd src/github.com/pawelhenek/go-by-example && go run ./multiple_return_values.go
    cd src/github.com/pawelhenek/go-by-example && go run ./mutexes.go
    cd src/github.com/pawelhenek/go-by-example && go run ./non-blocking_channel_operations.go
    cd src/github.com/pawelhenek/go-by-example && go run ./numbers-parsing.go
    #cd src/github.com/pawelhenek/go-by-example && go run ./panic.go
    cd src/github.com/pawelhenek/go-by-example && go run ./pointers.go
    cd src/github.com/pawelhenek/go-by-example && go run ./random-numbers.go
    cd src/github.com/pawelhenek/go-by-example && go run ./range.go
    cd src/github.com/pawelhenek/go-by-example && go run ./range_over_channels.go
    cd src/github.com/pawelhenek/go-by-example && go run ./rate_limiting.go
    #cd src/github.com/pawelhenek/go-by-example && go run ./reading-files.go
    cd src/github.com/pawelhenek/go-by-example && go run ./recursion.go
    cd src/github.com/pawelhenek/go-by-example && go run ./regular-expressions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./select.go
    cd src/github.com/pawelhenek/go-by-example && go run ./sha1-hashes.go
    #cd src/github.com/pawelhenek/go-by-example && go run ./singnals.go
    cd src/github.com/pawelhenek/go-by-example && go run ./slices.go
    cd src/github.com/pawelhenek/go-by-example && go run ./sorting.go
    cd src/github.com/pawelhenek/go-by-example && go run ./sorting-by-functions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./spawning-processes.go
    cd src/github.com/pawelhenek/go-by-example && go run ./stateful_goroutines.go
    cd src/github.com/pawelhenek/go-by-example && go run ./string-formatting.go
    cd src/github.com/pawelhenek/go-by-example && go run ./string-functions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./structs.go
    cd src/github.com/pawelhenek/go-by-example && go run ./switch.go
    cd src/github.com/pawelhenek/go-by-example && go run ./tickers.go
    cd src/github.com/pawelhenek/go-by-example && go run ./time.go
    cd src/github.com/pawelhenek/go-by-example && go run ./time-formatting-parsing.go
    cd src/github.com/pawelhenek/go-by-example && go run ./timeouts.go
    cd src/github.com/pawelhenek/go-by-example && go run ./timers.go
    cd src/github.com/pawelhenek/go-by-example && go run ./url-parsing.go
    cd src/github.com/pawelhenek/go-by-example && go run ./values.go
    cd src/github.com/pawelhenek/go-by-example && go run ./variables.go
    cd src/github.com/pawelhenek/go-by-example && go run ./variadic_functions.go
    cd src/github.com/pawelhenek/go-by-example && go run ./worker_pools.go
    cd src/github.com/pawelhenek/go-by-example && go run ./writting-files.go

run-learn-go-with-tests:
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./hello
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./integers
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./iteration
    cd src/github.com/pawelhenek/learn-go-with-tests && cd iteration && go test -bench ./
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./arrays_and_slices
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./structs_methods_interfaces
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./pointers_and_errors
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./maps
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./dependency-injection;
    cd src/github.com/pawelhenek/learn-go-with-tests && go test -v ./mocking

