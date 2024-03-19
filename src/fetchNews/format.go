package fetchNews

import (
	"newsSharing/src/guardian"
	"newsSharing/src/nyTimes"
	"strconv"
)

func FormatNews(guardianData *guardian.NewsData, nyTimesData *nyTimes.NewsData, page string) *ApiResponse {
	apiResponse := new(ApiResponse)
	apiResponse.Heading = "NEWS"
	apiResponse.PageNumber, _ = strconv.Atoi(page)
	if apiResponse.PageNumber == 0 {
		apiResponse.PageNumber = 1
	}
	apiResponse.TotalPages = guardianData.GetResponse().GetPages()
	apiResponse.Data = make([]*News, 0)

	for _, result := range guardianData.GetResponse().GetResults() {
		apiResponse.Data = append(apiResponse.Data, &News{
			Title:     result.GetWebTitle(),
			Url:       result.GetWebUrl(),
			Publisher: "Guardian",
		})
	}

	for _, doc := range nyTimesData.GetResponse().GetDocs() {
		apiResponse.Data = append(apiResponse.Data, &News{
			Title:     doc.GetHeadline().GetHeadLine(),
			Url:       doc.GetWebUrl(),
			Publisher: "NyTimes",
		})
	}

	if len(apiResponse.Data) == 0 {
		apiResponse.ErrorMessage = "Downstream servers down"
	}

	return apiResponse
}
