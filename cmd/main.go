package main

import (
	"os"

	"github.com/gardusig/guessserver/guess"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	server := guess.NewGuessServer()
	logrus.Debug("starting server...")
	err := server.Start()
	if err != nil {
		logrus.Error("failed to start guess server:", err)
		os.Exit(1)
	}
	logrus.Debug("started server")
}
