package logger

import "newsSharing/src/constants"

func GetLogger(loggerOutput string) ILogger {
	switch loggerOutput {
	case constants.STDOUT_LOGGING:
		return new(StdLogger)
	}

	return nil
}
