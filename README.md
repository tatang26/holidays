## Coding session: Go 1.24

Welcome to the coding session: Go 1.24! In this workshop, you'll be applying some of the new features introduced in Go 1.24 to experiment with them in a simple Leapkit application.

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

1. Let's use iterators to loop over the holidays array.
2. In the `internal/holiday/service.go` file, there is a function called `holidaysFor` that is in charge of retrieving holiday information for a specified country. You will need to update this function to hold results of previous holiday requests using a basic `map`.
3. The final task is to run your app using the official Leapkit tool `go.leapkit.dev/tools/dev`. To install it, we will need to install the `leapkit` tool by running `go get -tool go.leapkit.dev/tools/dev` and then, run `go tool dev`.
