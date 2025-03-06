## Coding session: Go 1.24

Welcome to the coding session: Go 1.24! In this workshop, you'll be applying some new features introduced in Go 1.24 to experiment with them in a simple Leapkit application.

## Prerequisites

- **Go 1.24** installed (Download Go)
- This repository
- To be willing to participate :D

### Running the application

**Initially**, you can run the application by executing the following command:
```sh
go run ./cmd/app
```

## Tasks

In the `internal/holiday/service.go` file, there is a function called `holidaysFor` that retrieves holiday information for a specified country. Our goal is to modify it to make it more efficient.

1. The function currently returns a list of holidays; to achieve this, you'll need to loop over the dates using an iterator to populate the holiday slice.
2. To prevent the function from making duplicate API requests, we will implement a caching mechanism using a simple `map` to store the results of previous requests.
3. Finally, we will change how we run the app. Instead of using the `go run` command, we will use the official Leapkit tool by downloading it from `go.leapkit.dev/tools/dev`. To install it, please run `go get -tool go.leapkit.dev/tools/dev`. Then, to run the app, use the command `go tool dev`.