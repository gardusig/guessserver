package internal

import "os"

const (
	guessServicePortEnvKey = "GUESS_SERVICE_GRPC_PORT"
	guessServiceHostEnvKey = "GUESS_SERVICE_GRPC_HOST"

	guessServiceDefaultPort = "50051"
	guessServiceDefaultHost = "localhost"

	LevelMinThreshold uint32 = 0
	LevelMaxThreshold uint32 = 1000

	GuessMinThreshold int64 = -4000000000000000000
	GuessMaxThreshold int64 = +4000000000000000000
)

var (
	GuessServicePort string
	GuessServiceHost string

	Equal   = "="
	Less    = "<"
	Greater = ">"
)

func init() {
	GuessServicePort = os.Getenv(guessServicePortEnvKey)
	if GuessServicePort == "" {
		GuessServicePort = guessServiceDefaultPort
	}
	GuessServiceHost = os.Getenv(guessServiceHostEnvKey)
	if GuessServiceHost == "" {
		GuessServiceHost = guessServiceDefaultHost
	}
}
