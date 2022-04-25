package utils

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	logger "github.com/sirupsen/logrus"
)

var ddClient *statsd.Client

const namespace = "{{APP_NAME}}"

func ConnectDatadog() {
	url := fmt.Sprintf("%s:8125", config.DataDogHost)
	var err error
	ddClient, err = statsd.New(url, statsd.WithNamespace(namespace))
	if err != nil {
		logger.WithError(err).Error("Datadog client could not be created")
	}
	logger.Infoln("Datadog client has been created successfully")
}

func RecordTime(msg string, tags []string) func() {
	start := time.Now()

	return func() {
		if err := ddClient.Timing(msg, time.Since(start), tags, 1); err != nil {
			logger.WithError(err).Error("could not measure elapsed time")
		}
	}
}
