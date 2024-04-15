# GuessServer

## Project Overview

GuessServer is the server-side application for the Guess service. It provides a gRPC interface for handling guess requests and managing game logic.

## How to Build and Run

To build and run GuessServer locally, follow these steps:

```bash
$ docker build . -t guessserver
$ docker run guessserver
```

## Configuration Options
GuessServer can be configured using environment variables:

`GUESS_SERVER_GRPC_PORT`: The port of the Guess server gRPC service (default: `50051`).
`GUESS_SERVER_GRPC_HOST`: The host of the Guess server gRPC service (default: `localhost`).

## Code Structure
The project directory structure is organized as follows:

```
guessserver/
├── cmd/
├── guess/
└── internal/
```

- `cmd/`: Contains the main entry point of the application.
- `guess/`: Contains the server logic for handling guess requests and game logic.
- `internal/`: Contains configuration constants and initialization logic.


## Dependencies

- `github.com/gardusig/guessproto`: Protobuf definitions for communication between the client and server.
