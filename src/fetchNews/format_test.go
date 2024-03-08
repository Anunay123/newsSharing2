package fetchNews

import (
	"newsSharing/src/guardian"
	"newsSharing/src/nyTimes"
	"testing"
)

func TestFormatNewsWDataFromBothAPIs(t *testing.T) {
	guardianData := new(guardian.NewsData)
	guardianData.Response = &guardian.Response{
		Pages:       10,
		CurrentPage: 1,
	}

	guardianData.Response.Results = make([]*guardian.Result, 0, 2)

	guardianData.Response.Results = append(guardianData.Response.Results, &guardian.Result{
		WebUrl:   "guardian-web-url-sample-1",
		WebTitle: "guardian-web-title-sample-1",
	}, &guardian.Result{
		WebUrl:   "guardian-web-url-sample-2",
		WebTitle: "guardian-web-title-sample-2",
	})

	nyTimesData := new(nyTimes.NewsData)
	nyTimesData.Response = &nyTimes.Response{
		Docs: make([]*nyTimes.Doc, 0, 3),
	}

	nyTimesData.Response.Docs = append(nyTimesData.Response.Docs, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-1",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-1",
		},
	}, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-2",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-2",
		},
	}, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-3",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-3",
		},
	})

	apiResponse := FormatNews(guardianData, nyTimesData, "1")

	expectedValueTotalNewsItems := 5
	actualValueTotalNewsItems := len(apiResponse.Data)

	if expectedValueTotalNewsItems != actualValueTotalNewsItems {
		t.Errorf("Expected Value of total news items: %d, Actual Value of total news items: %d", expectedValueTotalNewsItems, actualValueTotalNewsItems)
	}
}

func TestFormatNewsWDataFromNyTimes(t *testing.T) {
	nyTimesData := new(nyTimes.NewsData)
	nyTimesData.Response = &nyTimes.Response{
		Docs: make([]*nyTimes.Doc, 0, 3),
	}

	nyTimesData.Response.Docs = append(nyTimesData.Response.Docs, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-1",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-1",
		},
	}, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-2",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-2",
		},
	}, &nyTimes.Doc{
		WebUrl: "nyTimes-web-url-sample-3",
		Headline: &nyTimes.HeadLine{
			PrintHeadline: "nyTimes-web-title-sample-3",
		},
	})

	apiResponse := FormatNews(nil, nyTimesData, "3")

	expectedValueTotalNewsItems := 3
	actualValueTotalNewsItems := len(apiResponse.Data)

	if expectedValueTotalNewsItems != actualValueTotalNewsItems {
		t.Errorf("Expected Value of total news items: %d, Actual Value of total news items: %d", expectedValueTotalNewsItems, actualValueTotalNewsItems)
	}

	expectedValuePageNumber := 3
	actualValuePageNumber := apiResponse.PageNumber

	if expectedValuePageNumber != actualValuePageNumber {
		t.Errorf("Expected Value of page: %d, Actual Value of page: %d", expectedValuePageNumber, actualValuePageNumber)
	}
}

func TestFormatNewsWDataFromGuardian(t *testing.T) {
	guardianData := new(guardian.NewsData)
	guardianData.Response = &guardian.Response{
		Pages:       10,
		CurrentPage: 1,
	}

	guardianData.Response.Results = make([]*guardian.Result, 0, 2)

	guardianData.Response.Results = append(guardianData.Response.Results, &guardian.Result{
		WebUrl:   "guardian-web-url-sample-1",
		WebTitle: "guardian-web-title-sample-1",
	}, &guardian.Result{
		WebUrl:   "guardian-web-url-sample-2",
		WebTitle: "guardian-web-title-sample-2",
	})

	apiResponse := FormatNews(guardianData, nil, "1")

	expectedValueTotalNewsItems := 2
	actualValueTotalNewsItems := len(apiResponse.Data)

	if expectedValueTotalNewsItems != actualValueTotalNewsItems {
		t.Errorf("Expected Value of total news items: %d, Actual Value of total news items: %d", expectedValueTotalNewsItems, actualValueTotalNewsItems)
	}
}
