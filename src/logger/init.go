package logger

import "newsSharing/src/constants"

var (
	LoggerUtil ILogger
)

func DoInit() {
	LoggerUtil = GetLogger(constants.STDOUT_LOGGING)
}
