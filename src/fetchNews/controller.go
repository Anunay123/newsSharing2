package fetchNews

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsSharing/src/constants"
	"newsSharing/src/guardian"
	"newsSharing/src/logger"
	"newsSharing/src/network"
	"newsSharing/src/nyTimes"
)

func NewsController(ginContext *gin.Context) {

	queryString, page := ginContext.Query(constants.PARAM_SEARCHED_QUERY), ginContext.Query(constants.PARAM_SEARCHED_PAGE)

	if queryString == "" {
		logger.LoggerUtil.PrintLog(constants.LOG_LEVEL_ERROR, "No query provided")
		ginContext.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(fmt.Sprintf("{\"error\":\"%s\"}", "No query provided")))
		return
	}

	if page == "" {
		page = "1"
	}

	guardianChannel, nyTimesChannel := make(chan *guardian.NewsData, 1), make(chan *nyTimes.NewsData, 1)

	//async calls
	go guardian.GetNewsData(queryString, page, guardianChannel, network.CommonRestService)
	go nyTimes.GetNewsData(queryString, page, nyTimesChannel, network.CommonRestService)

	nyTimesData := <-nyTimesChannel
	guardianData := <-guardianChannel

	apiResponse := FormatNews(guardianData, nyTimesData, page)

	var (
		finalResult []byte
		err         error
	)

	if finalResult, err = json.Marshal(apiResponse); err != nil {
		logger.LoggerUtil.PrintLog(constants.LOG_LEVEL_ERROR, fmt.Sprintf("Error while unmarshalling final api response. Error: %s", err.Error()))
		ginContext.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
		return
	}

	ginContext.Data(http.StatusOK, "application/json; charset=utf-8", finalResult)
}
