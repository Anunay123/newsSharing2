package guardian

import (
	"github.com/afex/hystrix-go/hystrix"
	"newsSharing/src/config"
)

var (
	baseUrl        string
	hystrixCommand string
)

func DoInit() {

	baseUrl = config.Config.GetString("guardian.BASE_URL")
	hystrixCommand = config.Config.GetString("guardian.HYSTRIX.COMMAND")

	// Hystrix : to guard against service variability
	hystrix.ConfigureCommand(hystrixCommand, hystrix.CommandConfig{
		Timeout:               int(config.Config.GetFloat64("guardian.HYSTRIX.TIMEOUT_SEC")),
		MaxConcurrentRequests: config.Config.GetInt("guardian.HYSTRIX.MAX_CONC_REQUESTS"),
		ErrorPercentThreshold: config.Config.GetInt("guardian.HYSTRIX.err_percent_plus_threshold"),
	})
}
