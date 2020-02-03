package main

import (
	"os"

	"github.com/docker/docker/pkg/health"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stderr)
	if err := health.RunHealthChecks(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
