package nyTimes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsSharing/src/constants"
	"newsSharing/src/logger"
	"newsSharing/src/network"
)

func GetNewsData(query, page string, newsDataChannel chan *NewsData) {
	queryMap := map[string]string{
		"q":       query,
		"api-key": constants.API_KEY_NYTIMES,
		"page":    page,
	}

	responseBytes, err, statusCode := network.CallRestService(http.MethodGet, baseUrl, constants.JSON, constants.JSON,
		nil, queryMap, nil, nil, http.Client{})

	if err != nil {
		//error handling
		logger.LoggerUtil.PrintLog(constants.LOG_LEVEL_ERROR, fmt.Sprintf("Error fetching Guardian Data. Error: %s, StatusCode: %s", err.Error(), string(statusCode)))
		return
	}

	var newsData *NewsData

	if err := json.Unmarshal(responseBytes, &newsData); err != nil {
		//log unmarshal error
		logger.LoggerUtil.PrintLog(constants.LOG_LEVEL_ERROR, fmt.Sprintf("Error unmarshalling Guardian data. Error: %s", err.Error()))
		return
	}
	newsDataChannel <- newsData
}
