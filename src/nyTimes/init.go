package nyTimes

import (
	"github.com/afex/hystrix-go/hystrix"
	"newsSharing/src/config"
)

var (
	baseUrl        string
	hystrixCommand string
)

func DoInit() {

	baseUrl = config.Config.GetString("nyTimes.BASE_URL")
	hystrixCommand = config.Config.GetString("nyTimes.HYSTRIX.COMMAND")

	// Hystrix : to guard against service variability
	hystrix.ConfigureCommand(hystrixCommand, hystrix.CommandConfig{
		Timeout:               int(config.Config.GetFloat64("nyTimes.HYSTRIX.TIMEOUT_SEC") * 1000),
		MaxConcurrentRequests: config.Config.GetInt("nyTimes.HYSTRIX.MAX_CONC_REQUESTS"),
		ErrorPercentThreshold: config.Config.GetInt("nyTimes.HYSTRIX.err_percent_plus_threshold"),
	})
}
