package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury-core-go-template/src/api/config"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	logger "github.com/sirupsen/logrus"
)

func NewRelic() gin.HandlerFunc {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("{{APP_NAME}}"),
		newrelic.ConfigLicense(config.NewRelicKey),
		newrelic.ConfigInfoLogger(os.Stdout),
		func(config *newrelic.Config) {
			config.ErrorCollector.IgnoreStatusCodes = []int{http.StatusUnauthorized, http.StatusForbidden}
		},
	)
	if err != nil {
		logger.WithError(err).Fatal("Panic connecting with Newrelic")
		panic(err)
	}

	return nrgin.Middleware(app)
}
