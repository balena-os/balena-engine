package health

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type HealthCheck interface {
	Description() string
	Check(context.Context) error
}

var healthChecks = []HealthCheck{
	enginePing{},
	containerdHealth{},
	diskSpace{},
}

func RunHealthChecks() error {
	startTTotal := time.Now()
	logrus.Info("Start running health checks")
	defer func() {
		logrus.WithField("duration", time.Now().Sub(startTTotal)).
			Info("Done running health checks")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	for _, check := range healthChecks {
		startT := time.Now()
		err = check.Check(ctx)

		log := logrus.WithFields(logrus.Fields{
			"check": check.Description(),
			"duration": time.Now().Sub(startT),
		})
		if err == nil {
			log.Info("Success")
		} else {
			log.WithError(err).Error("Failure")
		}
	}
	return err
}

func NewEngineClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}
