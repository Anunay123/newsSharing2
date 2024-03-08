package fetchNews

import (
	"newsSharing/src/guardian"
	"newsSharing/src/nyTimes"
	"strconv"
)

func FormatNews(guardianData *guardian.NewsData, nyTimesData *nyTimes.NewsData, page string) *ApiResponse {
	apiResponse := new(ApiResponse)
	apiResponse.Heading = "NEWS"
	apiResponse.PageNumber = guardianData.GetResponse().GetCurrentPage()
	if apiResponse.PageNumber == 0 {
		apiResponse.PageNumber, _ = strconv.Atoi(page)
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

	return apiResponse
}
